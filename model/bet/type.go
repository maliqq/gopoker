package bet

import (
	"fmt"
)

type ForcedBet string
type ActiveBet string
type PassiveBet string
type AnyBet string

type Type interface {
	Value() AnyBet
}

type Bet struct {
	Amount float64
	Type
}

const (
	Ante       ForcedBet = "ante"
	BringIn    ForcedBet = "bring-in"
	SmallBlind ForcedBet = "small-blind"
	BigBlind   ForcedBet = "big-blind"
	GuestBlind ForcedBet = "guest-blind"
	Straddle   ForcedBet = "straddle"

	Raise ActiveBet = "raise"
	Call  ActiveBet = "call"

	Check PassiveBet = "check"
	Fold  PassiveBet = "fold"

	Discard  AnyBet = "discard"
	StandPat AnyBet = "stand-pat"

	Show AnyBet = "show"
	Muck AnyBet = "muck"
)

func (any AnyBet) Value() AnyBet {
	return any
}

func (forced ForcedBet) Value() AnyBet {
	return AnyBet(string(forced))
}

func (active ActiveBet) Value() AnyBet {
	return AnyBet(string(active))
}

func (passive PassiveBet) Value() AnyBet {
	return AnyBet(string(passive))
}

func (b Bet) String() string {
	if b.Amount != 0. {
		return fmt.Sprintf("%s %.2f", string(b.Type.Value()), b.Amount)
	}
	return string(b.Type.Value())
}

func (b Bet) IsActive() bool {
	switch b.Type.(type) {
	case ActiveBet:
		return true
	}

	return false
}

func (b Bet) IsForced() bool {
	switch b.Type.(type) {
	case ForcedBet:
		return true
	}

	return false
}

func NewBet(t Type, amount float64) *Bet {
	return &Bet{Type: t, Amount: amount}
}

func NewFold() *Bet {
	return &Bet{Type: Fold}
}

func NewRaise(amount float64) *Bet {
	return &Bet{Type: Raise, Amount: amount}
}

func NewCheck() *Bet {
	return &Bet{Type: Check}
}

func NewCall(amount float64) *Bet {
	return &Bet{Type: Call, Amount: amount}
}
