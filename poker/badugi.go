package poker

import (
	"errors"
)

import (
	"gopoker/poker/hand"
)

var (
	BadugiRanks = []rankFunc{
		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.BadugiOne, hc.isBadugiOne()
		},

		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.BadugiFour, hc.isBadugiFour()
		},

		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.BadugiThree, hc.isBadugiThree()
		},

		func(hc *handCards) (hand.Rank, *Hand) {
			return hand.BadugiTwo, hc.isBadugiTwo()
		},
	}
)

func (hc *handCards) isBadugiOne() *Hand {
	if len(hc.groupKind) == 1 {
		cards := hc.Cards

		return &Hand{
			Value: cards[0:1],
		}
	}

	if len(hc.groupSuit) == 1 {
		card := hc.groupSuit[0].Min(hc.Ordering)

		return &Hand{
			Value: Cards{*card},
		}
	}

	return nil
}

func (hc *handCards) isBadugiFour() *Hand {
	if len(hc.groupKind) == 4 && len(hc.groupSuit) == 4 {
		cards := hc.cardsHelper.Arrange()

		return &Hand{
			Value: *cards,
		}
	}

	return nil
}

func (hc *handCards) isBadugiThree() *Hand {
	paired, hasPaired := hc.paired[2]
	suited, hasSuited := hc.suited[2]

	if !hasPaired && !hasSuited {
		return nil
	}

	var a, b, c *Card = nil, nil, nil

	if len(paired) == 1 && len(suited) != 2 {
		cards := paired[0]

		a = &cards[0]

		diff := hc.Cards.Diff(&cards)

		for _, card := range diff {
			if a.kind != card.kind {
				n := card
				if b == nil {
					b = &n
				} else {
					c = &n
				}
			}
		}

		if b.suit == c.suit {
			return nil
		}

		if a.suit == b.suit || a.suit == c.suit {
			a = &cards[1]
		}

	} else if !hasPaired && len(suited) == 1 {
		a = suited[0].Min(hc.Ordering)

		for _, card := range hc.Cards.Diff(&suited[0]) {
			if a.suit != card.suit {
				n := card
				if b == nil {
					b = &n
				} else {
					c = &n
				}
			}
		}

		if b.kind == c.kind {
			return nil
		}

	} else {
		return nil
	}

	cards := Cards{*a, *b, *c}.Arrange(hc.Ordering)

	return &Hand{
		Value: cards,
	}
}

func (hc *handCards) isBadugiTwo() *Hand {
	var a, b *Card

	sets, hasPaired := hc.paired[3]
	suited, hasSuited := hc.suited[3]

	if hasPaired {
		cards := sets[0]

		diff := hc.Cards.Diff(&cards)

		b = &diff[0]

		for _, card := range cards {
			if b.suit != card.suit {
				n := card
				a = &n
				break
			}
		}

	} else if hasSuited {
		cards := suited[0]

		diff := hc.Cards.Diff(&cards)

		a = &diff[0]

		c := Cards{}
		for _, card := range cards {
			if a.kind != card.kind {
				c = append(c, card)
			}
		}

		b = c.Min(hc.Ordering)

	} else if len(hc.groupSuit) > 0 {
		cards := hc.groupSuit[0]

		a = cards.Min(hc.Ordering)

		c := Cards{}
		for _, card := range hc.Cards.Diff(&cards) {
			if a.suit != card.suit && a.kind != card.kind {
				c = append(c, card)
			}
		}

		b = c.Min(hc.Ordering)

	} else {
		cards := hc.groupKind[0]

		a = &cards[0]

		c := Cards{}
		for _, card := range hc.Cards.Diff(&cards) {
			if a.kind != card.kind {
				c = append(c, card)
			}
		}

		b = c.Min(hc.Ordering)
	}

	cards := Cards{*a, *b}.Arrange(hc.Ordering)

	return &Hand{
		Value: cards,
	}
}

func isBadugi(cards *Cards) (*Hand, error) {
	if len(*cards) != 4 {
		return nil, errors.New("4 cards required to detect badugi hand")
	}

	hc := NewHandCards(cards, AceLow, false)

	hand := hc.Detect(BadugiRanks)

	return hand, nil
}
