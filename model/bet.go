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

type BetRange struct {
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

func (newBet *Bet) Validate(seat *Seat, betRange BetRange) error {
	switch newBet.Type {
	case bet.Fold:
		// no error
	case bet.Check:
		if betRange.Call != seat.Bet {
			return fmt.Errorf("Can't check: need to call=%.2f", betRange.Call)
		}

	case bet.Call, bet.Raise:
		amount := newBet.Amount

		if amount > seat.Stack {
			return fmt.Errorf("Can't bet: got amount=%.2f, stack=%.2f", amount, seat.Stack)
		}

		if newBet.Type == bet.Call {
			return validateRange(amount, betRange.Call, betRange.Call, amount == seat.Stack)
		}

		if newBet.Type == bet.Raise {
			return validateRange(amount, betRange.Min, betRange.Max, amount == seat.Stack)
		}
	}

	return nil
}

func validateRange(amount float64, min float64, max float64, all_in bool) error {
	if max == 0. {
		return fmt.Errorf("Nothing to bet: got amount=%.2f", amount)
	}

	if amount > max {
		return fmt.Errorf("Bet invalid: got amount=%.2f, required max=%.2f", amount, max)
	}

	if amount < min && !all_in {
		return fmt.Errorf("Bet invalid: got amount=%.2f, required min=%.2f", amount, min)
	}

	return nil
}
