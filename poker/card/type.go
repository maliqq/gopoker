package card

import (
	"fmt"
)

import (
	"gopoker/util/console"
)

type Kind byte
type Suit byte
type Tuple struct {
	Kind
	Suit
}

const (
	Spade Suit = iota
	Heart
	Diamond
	Club
)

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
	Kinds = "23456789TJQKA"
	Suits = "shdc"

	KindsNum = len(Kinds)
	SuitsNum = len(Suits)
	CardsNum = KindsNum * SuitsNum
)

var (
	SuitsUnicode = []string{"♠", "♥", "♦", "♣"}

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

	Masks = []uint64{
		0x0001000000000000,
		0x0002000000000000,
		0x0004000000000000,
		0x0008000000000000,
		0x0010000000000000,
		0x0020000000000000,
		0x0040000000000000,
		0x0080000000000000,
		0x0100000000000000,
		0x0200000000000000,
		0x0400000000000000,
		0x0800000000000000,
		0x1000000000000000,
		0x0000000100000000,
		0x0000000200000000,
		0x0000000400000000,
		0x0000000800000000,
		0x0000001000000000,
		0x0000002000000000,
		0x0000004000000000,
		0x0000008000000000,
		0x0000010000000000,
		0x0000020000000000,
		0x0000040000000000,
		0x0000080000000000,
		0x0000100000000000,
		0x0000000000010000,
		0x0000000000020000,
		0x0000000000040000,
		0x0000000000080000,
		0x0000000000100000,
		0x0000000000200000,
		0x0000000000400000,
		0x0000000000800000,
		0x0000000001000000,
		0x0000000002000000,
		0x0000000004000000,
		0x0000000008000000,
		0x0000000010000000,
		0x0000000000000001,
		0x0000000000000002,
		0x0000000000000004,
		0x0000000000000008,
		0x0000000000000010,
		0x0000000000000020,
		0x0000000000000040,
		0x0000000000000080,
		0x0000000000000100,
		0x0000000000000200,
		0x0000000000000400,
		0x0000000000000800,
		0x0000000000001000,
	}
)

var (
	Colors = map[Suit]string{
		Spade:   console.YELLOW,
		Heart:   console.RED,
		Diamond: console.CYAN,
		Club:    console.GREEN,
	}
)

func (kind Kind) String() string {
	return string(Kinds[kind])
}

func (suit Suit) String() string {
	return string(Suits[suit])
}

func (suit Suit) UnicodeString() string {
	return SuitsUnicode[suit]
}

func AllKinds() []Kind {
	kinds := make([]Kind, KindsNum)

	for i := 0; i < 13; i++ {
		kinds[i] = Kind(i)
	}

	return kinds
}

func AllSuits() []Suit {
	suits := make([]Suit, SuitsNum)

	for i := 0; i < 4; i++ {
		suits[i] = Suit(i)
	}

	return suits
}

func AllTuples() []Tuple {
	cards := make([]Tuple, CardsNum)

	k := 0
	for _, kind := range AllKinds() {
		for _, suit := range AllSuits() {
			cards[k] = Tuple{kind, suit}
			k++
		}
	}

	return cards
}

func MakeKind(kind int) (Kind, error) {
	if kind < 0 || kind >= len(Kinds) {
		return 0, fmt.Errorf("invalid kind index %d", kind)
	}

	return Kind(kind), nil
}

func MakeSuit(suit int) (Suit, error) {
	if suit < 0 || suit >= len(Suits) {
		return 0, fmt.Errorf("invalid suit index %d", suit)
	}

	return Suit(suit), nil
}
