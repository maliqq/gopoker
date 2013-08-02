package poker

import (
	"errors"
	"fmt"
	"strings"
)

import (
	"gopoker/poker/card"
	"gopoker/util/console"
)

// Card - struct of kind and suit
type Card struct {
	kind card.Kind
	suit card.Suit
}

// String - card to string
func (c Card) String() string {
	return c.kind.String() + c.suit.String()
}

// KindName - e.g. "eight"
func (c Card) KindName() string {
	return c.kind.Name()
}

// KindTitle - e.g. "Eight"
func (c Card) KindTitle() string {
	return strings.Title(c.KindName())
}

// UnicodeString - e.g. "K♠"
func (c Card) UnicodeString() string {
	return c.kind.String() + c.suit.UnicodeString()
}

// Int - card to int
func (c Card) Int() int {
	return (int(c.kind) << 2) + int(c.suit)
}

// Byte - card to byte
func (c Card) Byte() byte {
	return (byte(c.kind) << 2) + byte(c.suit) + 1
}

// ConsoleString - colorified unicode string
func (c Card) ConsoleString() string {
	return console.Color(c.suit.Color(), c.UnicodeString())
}

// NewCard - new card from byte
func NewCard(i byte) (*Card, error) {
	if i <= 0 {
		return nil, nil
	}
	if int(i) > card.CardsNum {
		return nil, errors.New("invalid card")
	}

	i -= 1
	return &Card{card.Kind(i >> 2), card.Suit(i % 4)}, nil
}

// MakeCard - make card from kind and suit ints
func MakeCard(kind int, suit int) (*Card, error) {
	k, kindErr := card.MakeKind(kind)
	if kindErr != nil {
		return nil, kindErr
	}

	s, suitErr := card.MakeSuit(suit)
	if suitErr != nil {
		return nil, suitErr
	}

	return &Card{k, s}, nil
}

// ParseCard - parse card from plain string, e.g. "Kh"
func ParseCard(s string) (*Card, error) {
	if len(s) == 2 {
		p := strings.Split(s, "")

		kind := strings.Index(card.Kinds, strings.ToUpper(p[0]))
		suit := strings.Index(card.Suits, strings.ToLower(p[1]))

		card, err := MakeCard(kind, suit)
		if err != nil {
			return nil, err
		}

		return card, nil
	}

	return nil, fmt.Errorf("can't parse card %s", s)
}

func kindIndex(kind card.Kind, ord Ordering) int {
	switch ord {
	case AceHigh:
		return int(kind) + 1
	case AceLow:
		if Ace == kind {
			return 0
		}
		return int(kind) + 1
	}

	return -1
}

// Index - index of card for specified ordering
func (c Card) Index(ord Ordering) int {
	return kindIndex(c.kind, ord)
}

// Equal - check equality
func (c *Card) Equal(other *Card) bool {
	return (c.kind == other.kind) && (c.suit == other.suit)
}

// Compare - compare two cards; -1: less, 0: equal, 1: greater
func (c *Card) Compare(o *Card, ord Ordering) int {
	a, b := c.Index(ord), o.Index(ord)

	if a < b {
		return -1
	}
	if a == b {
		return 0
	}

	return 1
}
