package card

import (
	"fmt"
)

// Kind - 0..12 integer value of card
type Kind byte

// Kinds
const (
	Deuce Kind = iota // Deuce - "2"
	Three             // Three - "3"
	Four              // Four - "4"
	Five              // Five - "5"
	Six               // Six - "6"
	Seven             // Seven - "7"
	Eight             // Eight - "8"
	Nine              // Nine - "9"
	Ten               // Ten - "T"
	Jack              // Jack - "J"
	Queen             // Queen - "Q"
	King              // King - "K"
	Ace               // Ace - "A"
)

// Consts
const (
	Kinds    = "23456789TJQKA" // Kinds - all kinds
	KindsNum = len(Kinds)      // KindsNum - 13
)

var (
	kindNames = map[Kind]string{
		Deuce: "deuce",
		Three: "three",
		Four:  "four",
		Five:  "five",
		Six:   "six",
		Seven: "seven",
		Eight: "eight",
		Nine:  "nine",
		Ten:   "ten",
		Jack:  "jack",
		Queen: "queen",
		King:  "king",
		Ace:   "ace",
	}
)

// String - kind to string
func (kind Kind) String() string {
	return string(Kinds[kind])
}

// Name - kind name
func (kind Kind) Name() string {
	return kindNames[kind]
}

// AllKinds - list of all possible kinds
func AllKinds() []Kind {
	kinds := make([]Kind, KindsNum)

	for i := 0; i < 13; i++ {
		kinds[i] = Kind(i)
	}

	return kinds
}

// MakeKind - create card from int representation of card
func MakeKind(kind int) (Kind, error) {
	if kind < 0 || kind >= len(Kinds) {
		return 0, fmt.Errorf("invalid kind index %d", kind)
	}

	return Kind(kind), nil
}
