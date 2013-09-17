package card

import (
	"fmt"
)

import (
	"gopoker/util"
)

// Suit - 0..3 integer value of suit
type Suit byte

// Suits
const (
	Spade   Suit = iota // Spade - spades
	Heart               // Heart - heart
	Diamond             // Diamond - diamond
	Club                // Club - club
)

// Consts
const (
	Suits    = "shdc"     // Suits - all suits
	SuitsNum = len(Suits) // SuitsNum - 4
)

var (
	suitsUnicode = []string{"♠", "♥", "♦", "♣"}

	colors = map[Suit]string{
		Spade:   util.Yellow,
		Heart:   util.Red,
		Diamond: util.Cyan,
		Club:    util.Green,
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
