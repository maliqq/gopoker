package poker

import (
	"errors"
	"gopoker/poker/hand"
)

var (
	BadugiRanks = []rankFunc{
		func(pocket *PocketCards) (hand.Rank, *Hand) {
			return hand.BadugiOne, pocket.isBadugiOne()
		},

		func(pocket *PocketCards) (hand.Rank, *Hand) {
			return hand.BadugiFour, pocket.isBadugiFour()
		},

		func(pocket *PocketCards) (hand.Rank, *Hand) {
			return hand.BadugiThree, pocket.isBadugiThree()
		},

		func(pocket *PocketCards) (hand.Rank, *Hand) {
			return hand.BadugiTwo, pocket.isBadugiTwo()
		},
	}
)

func (pocket *PocketCards) isBadugiOne() *Hand {
	if len(pocket.groupKind) == 1 {
		cards := *pocket.Cards()

		return &Hand{
			Value: cards[0:1],
		}
	}

	if len(pocket.groupSuit) == 1 {
		card := pocket.groupSuit[0].Min(pocket.Ordering())

		return &Hand{
			Value: Cards{*card},
		}
	}

	return nil
}

func (pocket *PocketCards) isBadugiFour() *Hand {
	if len(pocket.groupKind) == 4 && len(pocket.groupSuit) == 4 {
		cards := ArrangeCards(pocket.Cards(), pocket.Ordering())

		return &Hand{
			Value: *cards,
		}
	}

	return nil
}

func (pocket *PocketCards) isBadugiThree() *Hand {
	paired, hasPaired := (*pocket.paired)[2]
	suited, hasSuited := (*pocket.suited)[2]

	if !hasPaired && !hasSuited {
		return nil
	}

	var a, b, c *Card = nil, nil, nil

	if len(paired) == 1 && len(suited) != 2 {
		cards := paired[0]

		a = &cards[0]

		diff := DiffCards(pocket.Cards(), &cards)

		for _, card := range *diff {
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
		a = suited[0].Min(pocket.Ordering())

		diff := DiffCards(pocket.Cards(), &suited[0])

		for _, card := range *diff {
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

	cards := ArrangeCards(&Cards{*a, *b, *c}, pocket.Ordering())

	return &Hand{
		Value: *cards,
	}
}

func (pocket *PocketCards) isBadugiTwo() *Hand {
	var a, b *Card

	sets, hasPaired := (*pocket.paired)[3]
	suited, hasSuited := (*pocket.suited)[3]

	if hasPaired {
		cards := sets[0]

		diff := DiffCards(pocket.Cards(), &cards)

		b = &(*diff)[0]

		for _, card := range cards {
			if b.suit != card.suit {
				n := card
				a = &n
				break
			}
		}

	} else if hasSuited {
		cards := suited[0]

		diff := DiffCards(pocket.Cards(), &cards)

		a = &(*diff)[0]

		c := Cards{}
		for _, card := range cards {
			if a.kind != card.kind {
				c = append(c, card)
			}
		}

		b = c.Min(pocket.Ordering())

	} else if len(pocket.groupSuit) > 0 {
		cards := pocket.groupSuit[0]

		a = cards.Min(pocket.Ordering())

		diff := DiffCards(pocket.Cards(), &cards)

		c := Cards{}
		for _, card := range *diff {
			if a.suit != card.suit && a.kind != card.kind {
				c = append(c, card)
			}
		}

		b = c.Min(pocket.Ordering())

	} else {
		cards := pocket.groupKind[0]

		a = &cards[0]

		diff := DiffCards(pocket.Cards(), &cards)

		c := Cards{}
		for _, card := range *diff {
			if a.kind != card.kind {
				c = append(c, card)
			}
		}

		b = c.Min(pocket.Ordering())
	}

	cards := ArrangeCards(&Cards{*a, *b}, pocket.Ordering())

	return &Hand{
		Value: *cards,
	}
}

func isBadugi(c *Cards) (*Hand, error) {
	if len(*c) != 4 {
		return nil, errors.New("4 cards required to detect badugi hand")
	}

	pocket := NewPocket(&OrderedCards{c, AceLow})

	hand := pocket.Detect(BadugiRanks)

	return hand, nil
}
