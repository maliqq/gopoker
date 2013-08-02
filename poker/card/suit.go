package card

import (
	"fmt"
)

import (
	"gopoker/util/console"
)

// Suit - 0..3 integer value of suit
type Suit byte

const (
	// Spade - spades
	Spade Suit = iota
	// Heart - heart
	Heart
	// Diamond - diamond
	Diamond
	// Club - club
	Club
)

const (
	// Suits - all suits
	Suits = "shdc"
	// SuitsNum - 4
	SuitsNum = len(Suits)
)

var (
	suitsUnicode = []string{"♠", "♥", "♦", "♣"}

	colors = map[Suit]string{
		Spade:   console.Yellow,
		Heart:   console.Red,
		Diamond: console.Cyan,
		Club:    console.Green,
	}
)

// String - suit to string, e.g. "s"
func (suit Suit) String() string {
	return string(Suits[suit])
}

// Color - console color of suit
func (suit Suit) Color() string {
	return colors[suit]
}

// UnicodeString - one of unicode chars for suit, e.g. "♠"
func (suit Suit) UnicodeString() string {
	return suitsUnicode[suit]
}

// AllSuits - all suits
func AllSuits() []Suit {
	suits := make([]Suit, SuitsNum)

	for i := 0; i < 4; i++ {
		suits[i] = Suit(i)
	}

	return suits
}

// MakeSuit - build suit from integer value
func MakeSuit(suit int) (Suit, error) {
	if suit < 0 || suit >= len(Suits) {
		return 0, fmt.Errorf("invalid suit index %d", suit)
	}

	return Suit(suit), nil
}
