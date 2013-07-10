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
	*model.Deal
	*model.Table
	model.Variation

	*protocol.Broadcast
	Receive chan *protocol.Message
	Control chan command.Type

	*GameRotation
	*Betting
	*Discarding
}

func (this *Play) String() string {
	return fmt.Sprintf("Game: %s\nTable: %s\n", this.Game(), this.Table)
}

// always returns current game
func (this *Play) Game() *model.Game {
	switch this.Variation.(type) {
	case *model.Game:
		return this.Variation.(*model.Game)
	case *model.Mix:
		return this.GameRotation.Current()
	default:
		fmt.Printf("got: %#v\n", this.Variation)
		panic("unknown game")
	}

	return nil
}

func (this *Play) RotateGame() {
	switch this.Variation.(type) {
	case *model.Game:
		// do nothing
	case *model.Mix:
		this.GameRotation.Rotate()
	default:
		fmt.Printf("got: %#v\n", this.Variation)
		panic("unknown game")
	}
}

func NewPlay(variation model.Variation, table *model.Table) *Play {
	play := &Play{
		Variation: variation,
		Table:     table,
		Broadcast: protocol.NewBroadcast(),
		Receive:   make(chan *protocol.Message),
		Control:   make(chan command.Type),
	}

	switch variation.(type) {
	case model.Game:
		// ok
	case model.Mix:
		play.GameRotation = NewGameRotation(variation.(*model.Mix), 0)
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

/*********************************
* Deals
*********************************/
func (this *Play) StartNewDeal() {
	this.Deal = model.NewDeal()
	this.Betting = NewBetting()
	if this.Game().Options.Discards {
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
			s.SetPlaying()
		}
	}
}

/*********************************
* Antes
*********************************/
func (this *Play) PostAntes() {
	stake := this.Game().Stake

	for _, pos := range this.Table.SeatsInPlay() {
		seat := this.Table.Seat(pos)

		newBet := this.Betting.ForceBet(pos, bet.Ante, stake)

		this.Betting.AddBet(seat, newBet)

		this.Broadcast.All <- protocol.NewAddBet(pos, newBet)
	}
}

/*********************************
* Blinds
*********************************/
func (this *Play) postSmallBlind(pos int) {
	stake := this.Game().Stake

	t := this.Table
	newBet := this.Betting.ForceBet(pos, bet.SmallBlind, stake)

	err := this.Betting.AddBet(t.Seat(pos), newBet)
	if err != nil {
		log.Fatalf("Error adding small blind for %d: %s", pos, err)
	}

	this.Broadcast.All <- protocol.NewAddBet(pos, newBet)
}

func (this *Play) postBigBlind(pos int) {
	stake := this.Game().Stake

	t := this.Table
	newBet := this.Betting.ForceBet(pos, bet.BigBlind, stake)

	err := this.Betting.AddBet(t.Seat(pos), newBet)
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

	this.Broadcast.One(seat.Player) <- this.Betting.RequireBet(minPos, seat, this.Game())
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

		this.Broadcast.One(seat.Player) <- betting.RequireBet(pos, seat, this.Game())

		select {
		case msg := <-betting.Receive:
			betting.Add(seat, msg)

		case <-time.After(time.Duration(DefaultTimer) * time.Second):
			fmt.Println("timeout!")
		}
	}
}

func (this *Play) ResetBetting() {
	betting := this.Betting

	betting.Reset()

	for _, pos := range this.Table.SeatsInPlay() {
		seat := this.Table.Seat(pos)

		seat.SetPlaying()
	}

	this.Broadcast.All <- protocol.NewPotSummary(betting.Pot)
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

		pocket := d.Pocket(player)

		if hand := d.Rank(pocket, ranking, withBoard); hand != nil {
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
