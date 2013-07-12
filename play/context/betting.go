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

type Betting struct {
	raiseCount int `json:"-"`
	BigBets    bool

	Pot *model.Pot

	Required *protocol.RequireBet

	Receive chan *protocol.Message
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

func (this *Betting) Current() int {
	return this.Required.Pos
}

func (this *Betting) Reset() {
	this.Required.Call, this.Required.Min, this.Required.Max, this.raiseCount = 0., 0., 0., 0
}

func (this *Betting) ForceBet(pos int, betType bet.Type, stake *model.Stake) *bet.Bet {
	amount := stake.Amount(betType)

	this.Required.Pos = pos
	this.Required.Call = amount

	return &bet.Bet{
		Type:   betType,
		Amount: amount,
	}
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


func (this *Betting) RequireBet(pos int, seat *model.Seat, game *model.Game, stake *model.Stake) *protocol.Message {
	this.Required.Pos = pos
	this.Required.Min, this.Required.Max = this.RaiseRange(seat, game, stake)

	return protocol.NewRequireBet(seat, this.Required)
}

func (this *Betting) ValidateBet(seat *model.Seat, newBet *bet.Bet) error {
	require := this.Required

	switch newBet.Type {
	case bet.Check:
		if require.Call != 0. {
			return fmt.Errorf("Can't check, need to call: %.2f", require.Call)
		}

	case bet.Call, bet.Raise:
		amount := newBet.Amount

		if amount > seat.Stack {
			return fmt.Errorf("Bet amount is greater than available stack: amount=%.2f stack=%.2f", amount, seat.Stack)
		}

		if newBet.Type == bet.Call && amount != require.Call - seat.Bet {
			return fmt.Errorf("Call mismatch: got amount=%.2f need to call=%.2f", amount, require.Call)
		}

		if newBet.Type == bet.Raise {
			if require.Max == 0. {
				return fmt.Errorf("Raise not allowed in current betting: got amount=%.2f", amount)
			}

			raiseAmount := require.Call - amount

			if raiseAmount > require.Max {
				return fmt.Errorf("Raise invalid: got amount=%.2f required max=%.2f", amount, require.Max)
			}

			if raiseAmount < require.Min && amount != seat.Stack {
				return fmt.Errorf("Raise invalid: got amount=%.2f required min=%.2f", amount, require.Min)
			}
		}
	}

	return nil
}

func (this *Betting) Call(amount float64) {
	this.Required.Call = amount
}

func (this *Betting) Raise(amount float64) {
	this.raiseCount++
	this.Required.Call += amount
}

func (this *Betting) AddBet(seat *model.Seat, newBet *bet.Bet) error {
	if newBet.Type == bet.Fold {
		seat.Fold()

		return nil
	}

	err := this.ValidateBet(seat, newBet)

	if err != nil {
		seat.Fold() // force fold

	} else {
		amount := newBet.Amount

		this.Pot.Add(seat.Player.Id, amount, amount == seat.Stack)

		if newBet.IsForced() {
			// ante, blinds
			this.Call(amount)
			seat.SetBet(amount)

		} else if newBet.IsActive() {
			// raise, call
			if newBet.Type == bet.Raise {
				this.Raise(amount)
			}
			seat.PutBet(amount)
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
