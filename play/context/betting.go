package context

import (
	"fmt"
	"log"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/game"
	"gopoker/protocol"
)

const (
	MaxRaises = 8 // TODO into game options
)

type Betting struct {
	raiseCount int
	BigBets    bool

	Pot *model.Pot

	Required *protocol.RequireBet

	Receive chan *protocol.Message `json:"-"`
}

func NewBetting() *Betting {
	return &Betting{
		Pot: model.NewPot(),

		Required: &protocol.RequireBet{},

		Receive: make(chan *protocol.Message),
	}
}

func (this *Betting) String() string {
	return fmt.Sprintf("Require %s %s raiseCount: %d bigBets: %t pot total: %.2f",
		this.Required,
		this.raiseCount,
		this.BigBets,
		this.Pot.Total(),
	)
}

func (this *Betting) Clear() {
	this.Required.Call, this.Required.Min, this.Required.Max, this.raiseCount = 0., 0., 0., 0
}

func (this *Betting) Current() int {
	return this.Required.Pos
}

func (this *Betting) RaiseRange(seat *model.Seat, g *model.Game, stake *model.Stake) (float64, float64) {
	_, bb := stake.Blinds()

	switch g.Limit {
	case game.NoLimit:
		return bb, seat.Stack

	case game.PotLimit:
		return bb, this.Pot.Total()

	case game.FixedLimit:
		if this.BigBets {
			return bb * 2, bb * 2
		}
		return bb, bb
	}

	return 0., 0.
}

func (this *Betting) ForceBet(pos int, betType bet.Type, stake *model.Stake) *model.Bet {
	amount := stake.Amount(betType)

	this.Required.Pos = pos
	this.Required.Call = amount

	return model.NewBet(betType, amount)
}

func (this *Betting) RequireBet(pos int, seat *model.Seat, game *model.Game, stake *model.Stake) *protocol.Message {
	this.Required.Pos = pos

	if this.raiseCount >= MaxRaises {
		this.Required.Min, this.Required.Max = 0., 0.
	} else {
		this.Required.Min, this.Required.Max = this.RaiseRange(seat, game, stake)
	}

	return protocol.NewRequireBet(seat, this.Required)
}

func (this *Betting) called(amount float64) {
	this.Required.Call = amount
	fmt.Printf("this.Required.Call=%.2f\n", amount)
}

func (this *Betting) raised(amount float64) {
	this.raiseCount++
	this.Required.Call += amount
	fmt.Printf("this.Required.Call=%.2f\n", amount)
}

func (this *Betting) AddBet(seat *model.Seat, newBet *model.Bet) error {
	if newBet.Type == bet.Fold {
		seat.Fold()

		return nil
	}

	err := newBet.Validate(seat, this.Required.RequireBet)

	if err != nil {
		seat.Fold() // force fold

	} else {
		amount := newBet.Amount

		this.Pot.Add(seat.Player.Id, amount, amount == seat.Stack)

		if newBet.IsForced() {
			// ante, blinds
			this.called(amount)
			seat.SetBet(amount)

		} else if newBet.IsActive() {
			// raise, call
			if newBet.Type == bet.Raise {
				this.raised(amount)
			}
			seat.AddBet(amount)
		}
	}

	return err
}

func (this *Betting) Add(seat *model.Seat, msg *protocol.Message) {
	newBet := msg.Payload.(protocol.AddBet).Bet

	log.Printf("[betting] Player %s %s\n", seat.Player, newBet.String())

	err := this.AddBet(seat, &newBet)

	if err != nil {
		log.Printf("[betting] %s", err)
	}
}
