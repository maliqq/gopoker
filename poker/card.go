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

type Card struct {
	kind card.Kind
	suit card.Suit
}

func (card Card) String() string {
	return card.kind.String() + card.suit.String()
}

func (card Card) UnicodeString() string {
	return card.kind.String() + card.suit.UnicodeString()
}

func (c Card) Int() int {
	return (int(c.kind) << 2) + int(c.suit)
}

func (c Card) Byte() byte {
	return (byte(c.kind) << 2) + byte(c.suit)
}

func (c Card) ConsoleString() string {
	return fmt.Sprintf("%s%s%s", card.Colors[c.suit], c.UnicodeString(), console.RESET)
}

func NewCard(i byte) (*Card, error) {
	if i < 0 || int(i) >= card.CardsNum {
		return nil, errors.New("invalid card")
	}

	return &Card{card.Kind(i >> 2), card.Suit(i % 4)}, nil
}

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

func (c Card) Index(ord Ordering) int {
	return kindIndex(c.kind, ord)
}

func (c Card) Equal(other Card) bool {
	return (c.kind == other.kind) && (c.suit == other.suit)
}
