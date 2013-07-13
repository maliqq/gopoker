package model

import (
	"fmt"
)

import (
	"gopoker/model/bet"
)

type Bet struct {
	bet.Type
	Amount float64
}

type RequireBet struct {
	Call float64
	Min  float64
	Max  float64
}

func (b Bet) String() string {
	if b.Amount != 0. {
		return fmt.Sprintf("%s %.2f", string(b.Type.Value()), b.Amount)
	}
	return string(b.Type.Value())
}

func (b Bet) IsActive() bool {
	switch b.Type.(type) {
	case bet.ActiveBet:
		return true
	}

	return false
}

func (b Bet) IsForced() bool {
	switch b.Type.(type) {
	case bet.ForcedBet:
		return true
	}

	return false
}

func NewBet(t bet.Type, amount float64) *Bet {
	return &Bet{Type: t, Amount: amount}
}

func NewFold() *Bet {
	return &Bet{Type: bet.Fold}
}

func NewRaise(amount float64) *Bet {
	return &Bet{Type: bet.Raise, Amount: amount}
}

func NewCheck() *Bet {
	return &Bet{Type: bet.Check}
}

func NewCall(amount float64) *Bet {
	return &Bet{Type: bet.Call, Amount: amount}
}

func (newBet *Bet) Validate(seat *Seat, required *RequireBet) error {
	switch newBet.Type {
	case bet.Check:
		if required.Call != 0. {
			return fmt.Errorf("Can't check, need to call: %.2f", required.Call)
		}

	case bet.Call, bet.Raise:
		amount := newBet.Amount

		if amount > seat.Stack {
			return fmt.Errorf("Bet amount is greater than available stack: amount=%.2f stack=%.2f", amount, seat.Stack)
		}

		if newBet.Type == bet.Call && amount != required.Call-seat.Bet {
			return fmt.Errorf("Call mismatch: got amount=%.2f need to call=%.2f", amount, required.Call)
		}

		if newBet.Type == bet.Raise {
			if required.Max == 0. {
				return fmt.Errorf("Raise not allowed in current betting: got amount=%.2f", amount)
			}

			raiseAmount := required.Call - amount

			if raiseAmount > required.Max {
				return fmt.Errorf("Raise invalid: got amount=%.2f required max=%.2f", amount, required.Max)
			}

			if raiseAmount < required.Min && amount != seat.Stack {
				return fmt.Errorf("Raise invalid: got amount=%.2f required min=%.2f", amount, required.Min)
			}
		}
	}

	return nil
}
