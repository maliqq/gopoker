package context

import (
	"fmt"
	"log"
	"time"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/deal"
	"gopoker/model/seat"
	"gopoker/play/command"
	"gopoker/poker"
	"gopoker/poker/ranking"
	"gopoker/protocol"
	"gopoker/util/console"
)

/*********************************
* Context
*********************************/
type Play struct {
	// dealt cards context
	Deal *model.Deal

	// mixed or limited game
	Game          *model.Game
	Mix           *model.Mix
	*GameRotation `json:"-"`

	// betting price
	Stake *model.Stake

	// players seating context
	Table *model.Table

	// players action context
	*Betting    `json:"-"`
	*Discarding `json:"-"`

	// broadcast protocol messages
	Broadcast *protocol.Broadcast `json:"-"`

	// receive protocol messages
	Receive chan *protocol.Message `json:"-"`

	// manage play
	Control chan command.Type `json:"-"`
}

func (this *Play) String() string {
	return fmt.Sprintf("Game: %s\nTable: %s\n", this.Game, this.Table)
}

func (this *Play) RotateGame() {
	if this.Mix == nil {
		return
	}

	if nextGame := this.GameRotation.Next(); nextGame != nil {
		this.Game = nextGame
		this.Broadcast.All <- protocol.NewChangeGame(nextGame)
	}
}

func NewPlay(variation model.Variation, stake *model.Stake, table *model.Table) *Play {
	play := &Play{
		Table:     table,
		Stake:     stake,
		Broadcast: protocol.NewBroadcast(),
		Receive:   make(chan *protocol.Message),
		Control:   make(chan command.Type),
	}

	if variation.IsMixed() {
		mix := variation.(*model.Mix)
		play.Mix = mix
		play.GameRotation = NewGameRotation(mix, 1)
		play.Game = play.GameRotation.Current()
	} else {
		play.Game = variation.(*model.Game)
	}

	go play.receive()

	return play
}

func (this *Play) receive() {
	for {
		msg := <-this.Receive
		log.Printf(console.Color(console.YELLOW, msg.String()))

		switch msg.Payload.(type) {
		case protocol.JoinTable:
			join := msg.Payload.(protocol.JoinTable)
			this.Table.AddPlayer(join.Player, join.Pos, join.Amount)

			//this.Broadcast.Except(join.Player) <- join

		case protocol.LeaveTable:
			leave := msg.Payload.(protocol.LeaveTable)
			this.Table.RemovePlayer(leave.Player)

			//this.Broadcast.Except(join.Player) <- leave

			// TODO: fold & autoplay

		case protocol.SitOut:
			sitOut := msg.Payload.(protocol.SitOut)
			this.Table.Seat(sitOut.Pos).State = seat.Idle

			// TODO: fold

		case protocol.ComeBack:
			comeBack := msg.Payload.(protocol.ComeBack)
			this.Table.Seat(comeBack.Pos).State = seat.Ready

		case protocol.AddBet:
			this.Betting.Receive <- msg

		case protocol.DiscardCards:
			this.Discarding.Receive <- msg
		}
	}
}

func (this *Play) Start() {

}

func (this *Play) Pause() {

}

func (this *Play) Resume() {

}

func (this *Play) Close() {

}

/*********************************
* Deals
*********************************/
func (this *Play) StartNewDeal() {
	this.Deal = model.NewDeal()
	this.Betting = NewBetting()
	if this.Game.Discards {
		this.Discarding = NewDiscarding(this.Deal)
	}
}

func (this *Play) ScheduleNextDeal() {
	this.Control <- command.NextDeal
}

/*********************************
* Seats
*********************************/
func (this *Play) ResetSeats() {
	for _, s := range this.Table.Seats {
		switch s.State {
		case seat.Ready, seat.Play:
			s.Play()
		}
	}
}

/*********************************
* Antes
*********************************/
func (this *Play) PostAntes() {
	for _, pos := range this.Table.SeatsInPlay() {
		seat := this.Table.Seat(pos)

		newBet := this.ForceBet(pos, bet.Ante, this.Stake)

		this.AddBet(seat, newBet)

		this.Broadcast.All <- protocol.NewAddBet(pos, newBet)
	}
}

/*********************************
* Blinds
*********************************/
func (this *Play) postSmallBlind(pos int) {
	t := this.Table
	newBet := this.ForceBet(pos, bet.SmallBlind, this.Stake)

	err := this.AddBet(t.Seat(pos), newBet)
	if err != nil {
		log.Fatalf("Error adding small blind for %d: %s", pos, err)
	}

	this.Broadcast.All <- protocol.NewAddBet(pos, newBet)
}

func (this *Play) postBigBlind(pos int) {
	t := this.Table
	newBet := this.ForceBet(pos, bet.BigBlind, this.Stake)

	err := this.AddBet(t.Seat(pos), newBet)
	if err != nil {
		log.Fatalf("Error adding big blind for %d: %s", pos, err)
	}

	this.Broadcast.All <- protocol.NewAddBet(pos, newBet)
}

func (this *Play) PostBlinds() {
	t := this.Table

	active := t.Seats.From(t.Button).Active()
	waiting := t.Seats.From(t.Button).Waiting()

	if len(active)+len(waiting) < 2 {
		log.Println("[this.stage.blinds] none waiting")

		return
	}
	//headsUp := len(active) == 2 && len(waiting) == 0 || len(active) == 1 && len(waiting) == 1

	sb := active[0]
	this.postSmallBlind(sb)

	bb := active[1]
	this.postBigBlind(bb)
}

/*********************************
* Button
*********************************/
func (this *Play) SetButton(pos int) {
	this.Table.SetButton(pos)

	this.Broadcast.All <- protocol.NewMoveButton(pos)
}

func (this *Play) MoveButton() {
	this.Table.MoveButton()

	this.Broadcast.All <- protocol.NewMoveButton(this.Table.Button)
}

/*********************************
* Bring in
*********************************/
func (this *Play) BringIn() {
	minPos := 0
	var card poker.Card

	for _, pos := range this.Table.SeatsInPlay() {
		s := this.Table.Seat(pos)

		pocketCards := *this.Deal.Pocket(s.Player)

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

	this.SetButton(minPos)

	seat := this.Table.Seat(minPos)

	this.Broadcast.One(seat.Player) <- this.Betting.RequireBet(minPos, seat, this.Game, this.Stake)
}

/*********************************
* Dealing
*********************************/
func (this *Play) DealHole(cardsNum int) {
	for _, pos := range this.Table.SeatsInPlay() {
		player := this.Table.Player(pos)

		cards := this.Deal.DealPocket(player, cardsNum)

		this.Broadcast.One(player) <- protocol.NewDealPocket(pos, cards, deal.Hole)
	}
}

func (this *Play) DealDoor(cardsNum int) {
	for _, pos := range this.Table.SeatsInPlay() {
		player := this.Table.Player(pos)

		cards := this.Deal.DealPocket(player, cardsNum)

		this.Broadcast.All <- protocol.NewDealPocket(pos, cards, deal.Door)
	}
}

func (this *Play) DealBoard(cardsNum int) {
	cards := this.Deal.DealBoard(cardsNum)

	this.Broadcast.All <- protocol.NewDealShared(cards, deal.Board)
}

/*********************************
* Betting
*********************************/
const (
	DefaultTimer = 30
)

func (this *Play) StartBettingRound() {
	betting := this.Betting

	for _, pos := range this.Table.Seats.From(betting.Current()).InPlay() {
		seat := this.Table.Seat(pos)

		this.Broadcast.One(seat.Player) <- betting.RequireBet(pos, seat, this.Game, this.Stake)

		select {
		case msg := <-betting.Receive:
			betting.Add(seat, msg)

		case <-time.After(time.Duration(DefaultTimer) * time.Second):
			fmt.Println("timeout!")
		}
	}
}

func (this *Play) ResetBetting() {
	this.Betting.Clear()

	for _, pos := range this.Table.SeatsInPlay() {
		seat := this.Table.Seat(pos)
		seat.Play()
	}

	this.Broadcast.All <- protocol.NewPotSummary(this.Pot)
}

/*********************************
* Discarding
*********************************/
func (this *Play) StartDiscardingRound() {
	discarding := this.Discarding

	for _, pos := range this.Table.SeatsFromButton().InPlay() {
		seat := this.Table.Seat(pos)

		this.Broadcast.One(seat.Player) <- discarding.RequireDiscard(pos)

		select {
		case msg := <-discarding.Receive:
			player, cards := discarding.Add(seat, msg)
			this.discard(player, cards)

		case <-time.After(time.Duration(DefaultTimer) * time.Second):
			fmt.Println("timeout!")
		}
	}
}

func (this *Play) discard(p *model.Player, cards *poker.Cards) {
	pos, _ := this.Table.Pos(p)

	cardsNum := len(*cards)

	this.Broadcast.All <- protocol.NewDiscarded(pos, cardsNum)

	if cardsNum > 0 {
		newCards := this.Deal.Discard(p, cards)

		this.Broadcast.One(p) <- protocol.NewDealPocket(pos, newCards, deal.Discard)
	}
}

/*********************************
* Showdown
*********************************/
type ShowdownHands map[model.Id]*poker.Hand

func (this *Play) ShowHands(ranking ranking.Type, withBoard bool) *ShowdownHands {
	d := this.Deal

	hands := ShowdownHands{}

	for _, pos := range this.Table.SeatsInPot() {
		player := this.Table.Player(pos)
		if pocket, hand := d.Rank(player, ranking, withBoard); hand != nil {
			hands[player.Id] = hand

			this.Broadcast.All <- protocol.NewShowHand(pos, pocket, hand)
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

func (this *Play) Winners(highHands *ShowdownHands, lowHands *ShowdownHands) {
	pot := this.Betting.Pot

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
			this.Broadcast.All <- protocol.NewWinner(winnerLow, total/2.)
			this.Broadcast.All <- protocol.NewWinner(winnerHigh, total/2.)
		} else {
			var exclusiveWinner model.Id

			if hi {
				exclusiveWinner = winnerHigh
			} else {
				exclusiveWinner = winnerLow
			}

			this.Broadcast.All <- protocol.NewWinner(exclusiveWinner, total)
		}
	}
}
