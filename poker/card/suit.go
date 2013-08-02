package card

import (
	"fmt"
)

import (
	"gopoker/util/console"
)

type Suit byte

const (
	Spade Suit = iota
	Heart
	Diamond
	Club
)

const (
	Suits    = "shdc"
	SuitsNum = len(Suits)
)

var (
	SuitsUnicode = []string{"♠", "♥", "♦", "♣"}

	Colors = map[Suit]string{
		Spade:   console.Yellow,
		Heart:   console.Red,
		Diamond: console.Cyan,
		Club:    console.Green,
	}
)

func (suit Suit) String() string {
	return string(Suits[suit])
}

func (suit Suit) UnicodeString() string {
	return SuitsUnicode[suit]
}

func AllSuits() []Suit {
	suits := make([]Suit, SuitsNum)

	for i := 0; i < 4; i++ {
		suits[i] = Suit(i)
	}

	return suits
}

func MakeSuit(suit int) (Suit, error) {
	if suit < 0 || suit >= len(Suits) {
		return 0, fmt.Errorf("invalid suit index %d", suit)
	}

	return Suit(suit), nil
}
