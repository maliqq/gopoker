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

	Masks = AllMasks()
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

func AllMasks() []uint64 {
	masks := make([]uint64, CardsNum)
	i := 0;
	for suit := SuitsNum-1; suit >= 0; suit-- {
		for kind := 0; kind < KindsNum; kind++ {
			masks[i] = uint64((1 << uint(kind) << uint(1 << 4 * suit)))
			i++
		}
	}
	return masks
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
