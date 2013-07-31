package card

import (
	"fmt"
)

type Kind byte

const (
	Deuce Kind = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	Kinds    = "23456789TJQKA"
	KindsNum = len(Kinds)
)

var (
	KindNames = map[Kind]string{
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

func (kind Kind) String() string {
	return string(Kinds[kind])
}

func AllKinds() []Kind {
	kinds := make([]Kind, KindsNum)

	for i := 0; i < 13; i++ {
		kinds[i] = Kind(i)
	}

	return kinds
}

func MakeKind(kind int) (Kind, error) {
	if kind < 0 || kind >= len(Kinds) {
		return 0, fmt.Errorf("invalid kind index %d", kind)
	}

	return Kind(kind), nil
}
