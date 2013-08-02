package card

import (
	"fmt"
)

// Kind - 0..12 integer value of card
type Kind byte

const (
	// Deuce - "2"
	Deuce Kind = iota
	// Three - "3"
	Three
	// Four - "4"
	Four
	// Five - "5"
	Five
	// Six - "6"
	Six
	// Seven - "7"
	Seven
	// Eight - "8"
	Eight
	// Nine - "9"
	Nine
	// Ten - "T"
	Ten
	// Jack - "J"
	Jack
	// Queen - "Q"
	Queen
	// King - "K"
	King
	// Ace - "A"
	Ace
)

const (
	// Kinds - all kinds
	Kinds = "23456789TJQKA"
	// KindsNum - 13
	KindsNum = len(Kinds)
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
