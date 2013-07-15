package model

import (
	"fmt"
)

import (
	"gopoker/poker"
)

type Dealer struct {
	Deck      poker.Cards
	burned    poker.Cards
	discarded poker.Cards
	shared    poker.Cards
	dealt     poker.Cards
}

func NewDealer() *Dealer {
	return &Dealer{
		Deck:      *poker.NewDeck(),
		burned:    poker.Cards{},
		discarded: poker.Cards{},
		shared:    poker.Cards{},
		dealt:     poker.Cards{},
	}
}

func (dealer *Dealer) give(n int) (*poker.Cards, error) {
	if l := len(dealer.Deck); l < n {
		return nil, fmt.Errorf("can't deal %d cards in deck with %d cards", n, l)
	}

	cards := dealer.Deck[0:n]

	dealer.Deck = dealer.Deck.Diff(&cards)

	return &cards, nil
}

func (dealer *Dealer) burn(n int) {
	cards, _ := dealer.give(n)

	dealer.burned = append(dealer.burned, *cards...)
}

func (dealer *Dealer) Burn(cards *poker.Cards) {
	dealer.Deck = dealer.Deck.Diff(cards)
	dealer.burned = append(dealer.burned, *cards...)
}

func (dealer *Dealer) Discard(cards *poker.Cards) *poker.Cards {
	n := len(*cards)
	if n > len(dealer.Deck) {
		dealer.reshuffle()
	}

	dealt, _ := dealer.give(n)

	dealer.burned = append(dealer.burned, *cards...)

	dealer.dealt = append(dealer.dealt.Diff(cards), *dealt...)

	return dealt
}

func (dealer *Dealer) Share(n int) *poker.Cards {
	dealer.burn(1)

	cards, _ := dealer.give(n)

	dealer.shared = append(dealer.shared, *cards...)

	return cards
}

func (dealer *Dealer) Deal(n int) *poker.Cards {
	cards, _ := dealer.give(n)

	dealer.dealt = append(dealer.dealt, *cards...)

	return cards
}

func (dealer *Dealer) reshuffle() {
	newDeck := append(dealer.Deck, dealer.burned...)

	dealer.Deck = *newDeck.Shuffle()
	dealer.burned = poker.Cards{}
}

func (dealer *Dealer) String() string {
	return fmt.Sprintf("dealer: deck=%s dealt=%s burned=%s", dealer.Deck, dealer.dealt, dealer.burned)
}
