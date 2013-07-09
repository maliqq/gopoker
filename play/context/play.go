package context

import (
	"fmt"
	"log"
	"time"
)

import (
	"gopoker/poker"
	"gopoker/poker/ranking"
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/seat"
	"gopoker/model/deal"
	"gopoker/play/command"
	"gopoker/protocol"
	"gopoker/util/console"
)

type Play struct {
	*model.Deal
	*model.Table
	*model.Game

	*protocol.Broadcast
	Receive chan *protocol.Message
	Control chan command.Type

	*Betting
	*Discarding
}

func (this *Play) String() string {
	return fmt.Sprintf("Game: %s\nTable: %s\n", this.Game, this.Table)
}

func NewPlay(game *model.Game, table *model.Table) *Play {
	play := &Play{
		Game:      game,
		Table:     table,
		Broadcast: protocol.NewBroadcast(),
		Receive:   make(chan *protocol.Message),
		Control:   make(chan command.Type),
	}

	go play.receive()

	return play
}

func (this *Play) NextDeal() {
	this.Deal = model.NewDeal()
	this.Betting = NewBetting()
	if this.Game.Options.Discards {
		this.Discarding = NewDiscarding()
	}
}

func (this *Play) receive() {
	for {
		msg := <-this.Receive
		log.Printf(console.Color(console.YELLOW, msg.String()))

		switch msg.Payload.(type) {
		case protocol.JoinTable:
			join := msg.Payload.(protocol.JoinTable)

			this.Table.AddPlayer(join.Player, join.Pos, join.Amount)

		case protocol.LeaveTable:
			leave := msg.Payload.(protocol.LeaveTable)

			this.Table.RemovePlayer(leave.Player)

		case protocol.AddBet:
			this.Betting.Receive <- msg

		case protocol.DiscardCards:
			this.Discarding.Receive <- msg
		}
	}
}

func (play *Play) ResetSeats() {
	for _, s := range play.Table.Seats {
		switch s.State {
		case seat.Ready, seat.Play:
			s.SetPlaying()
		}
	}
}

func (play *Play) PostAntes() {
	stake := play.Game.Stake
	
	for _, pos := range play.Table.SeatsInPlay() {
		seat := play.Table.Seat(pos)

		newBet := play.Betting.ForceBet(pos, bet.Ante, stake)

		play.Betting.AddBet(seat, newBet)

		play.Broadcast.All <- protocol.NewAddBet(pos, newBet)
	}
}

func (play *Play) postSmallBlind(pos int) {
	stake := play.Game.Stake

	t := play.Table
	newBet := play.Betting.ForceBet(pos, bet.SmallBlind, stake)

	err := play.Betting.AddBet(t.Seat(pos), newBet)
	if err != nil {
		log.Fatalf("Error adding small blind for %d: %s", pos, err)
	}

	play.Broadcast.All <- protocol.NewAddBet(pos, newBet)
}

func (play *Play) postBigBlind(pos int) {
	stake := play.Game.Stake

	t := play.Table
	newBet := play.Betting.ForceBet(pos, bet.BigBlind, stake)

	err := play.Betting.AddBet(t.Seat(pos), newBet)
	if err != nil {
		log.Fatalf("Error adding big blind for %d: %s", pos, err)
	}

	play.Broadcast.All <- protocol.NewAddBet(pos, newBet)
}

func (play *Play) PostBlinds() {
	t := play.Table

	active := t.Seats.From(t.Button).Active()
	waiting := t.Seats.From(t.Button).Waiting()

	if len(active)+len(waiting) < 2 {
		log.Println("[play.stage.blinds] none waiting")

		return
	}
	//headsUp := len(active) == 2 && len(waiting) == 0 || len(active) == 1 && len(waiting) == 1

	sb := active[0]
	play.postSmallBlind(sb)

	bb := active[1]
	play.postBigBlind(bb)
}

func (play *Play) SetButton(pos int) {
	play.Table.SetButton(pos)

	play.Broadcast.All <- protocol.NewMoveButton(pos)
}

func (play *Play) MoveButton() {
	play.Table.MoveButton()

	play.Broadcast.All <- protocol.NewMoveButton(play.Table.Button)
}

func (play *Play) BringIn() {
	minPos := 0
	var card poker.Card

	for _, pos := range play.Table.SeatsInPlay() {
		s := play.Table.Seat(pos)

		pocketCards := *play.Deal.Pocket(s.Player)

		lastCard := pocketCards[len(pocketCards)-1]
		if pos == 0 {
			card = lastCard
		} else {
			if lastCard.Compare(card, poker.AceHigh) > 0 {
				card = lastCard
				minPos = pos
			}
		}
	}

	play.SetButton(minPos)

	seat := play.Table.Seat(minPos)

	play.Broadcast.One(seat.Player) <- play.Betting.RequireBet(minPos, seat, play.Game)
}

func (play *Play) DealHole(cardsNum int) {
	for _, pos := range play.Table.SeatsInPlay() {
		player := play.Table.Player(pos)

		cards := play.Deal.DealPocket(player, cardsNum)

		play.Broadcast.One(player) <- protocol.NewDealPocket(pos, cards, deal.Hole)
	}
}

func (play *Play) DealDoor(cardsNum int) {
	for _, pos := range play.Table.SeatsInPlay() {
		player := play.Table.Player(pos)

		cards := play.Deal.DealPocket(player, cardsNum)

		play.Broadcast.All <- protocol.NewDealPocket(pos, cards, deal.Door)
	}
}

func (play *Play) DealBoard(cardsNum int) {
	cards := play.Deal.DealBoard(cardsNum)

	play.Broadcast.All <- protocol.NewDealShared(cards, deal.Board)
}

const (
	DefaultTimer = 30
)

func (play *Play) StartBettingRound() {
	betting := play.Betting

	for _, pos := range play.Table.Seats.From(betting.Current()).InPlay() {
		seat := play.Table.Seat(pos)

		play.Broadcast.One(seat.Player) <- betting.RequireBet(pos, seat, play.Game)

		select {
		case msg := <-betting.Receive:
			betting.Add(seat, msg)

		case <-time.After(time.Duration(DefaultTimer) * time.Second):
			fmt.Println("timeout!")
		}
	}
}

func (play *Play) discard(p *model.Player, cards *poker.Cards) {
	pos, _ := play.Table.Pos(p)

	cardsNum := len(*cards)

	play.Broadcast.All <- protocol.NewDiscarded(pos, cardsNum)

	if cardsNum > 0 {
		newCards := play.Deal.Discard(p, cards)

		play.Broadcast.One(p) <- protocol.NewDealPocket(pos, newCards, deal.Discard)
	}
}

func (play *Play) ResetBets() {
	betting := play.Betting

	betting.Reset()

	for _, pos := range play.Table.SeatsInPlay() {
		seat := play.Table.Seat(pos)

		seat.SetPlaying()
	}

	play.Broadcast.All <- protocol.NewPotSummary(betting.Pot)
}

type ShowdownHands map[model.Id]*poker.Hand

func (play *Play) ShowHands(ranking ranking.Type, withBoard bool) *ShowdownHands {
	d := play.Deal

	hands := ShowdownHands{}

	for _, pos := range play.Table.SeatsInPot() {
		player := play.Table.Player(pos)

		pocket := d.Pocket(player)

		if hand := d.Rank(pocket, ranking, withBoard); hand != nil {
			hands[player.Id] = hand

			play.Broadcast.All <- protocol.NewShowHand(pos, pocket, hand)
		}
	}

	return &hands
}

func best(sidePot *model.SidePot, hands *ShowdownHands) (model.Id, *poker.Hand) {
  var winner model.Id
  var best *poker.Hand

  for member, _ := range sidePot.Members {
    hand, hasHand := (*hands)[member]

    if hasHand && (best == nil || hand.Compare(best) > 0) {
      winner = member
      best = hand
    }
  }

  return winner, best
}

func (play *Play) Winners(highHands *ShowdownHands, lowHands *ShowdownHands) {
  pot := play.Betting.Pot

  hi := highHands != nil
  lo := lowHands != nil
  split := hi && lo

  for _, sidePot := range pot.SidePots() {
    total := sidePot.Total()

    var winnerLow, winnerHigh model.Id
    var bestLow *poker.Hand

    if lo {
      winnerLow, bestLow = best(sidePot, lowHands)
    }

    if hi {
      winnerHigh, _ = best(sidePot, highHands)
    }

    if split && bestLow != nil {
      play.Broadcast.All <- protocol.NewWinner(winnerLow, total/2.)
      play.Broadcast.All <- protocol.NewWinner(winnerHigh, total/2.)
    } else {
      var exclusiveWinner model.Id

      if hi {
        exclusiveWinner = winnerHigh
      } else {
        exclusiveWinner = winnerLow
      }

      play.Broadcast.All <- protocol.NewWinner(exclusiveWinner, total)
    }
  }
}
