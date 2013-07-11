package poker

import (
	"fmt"
)

import (
	"gopoker/poker/hand"
)

type handCards struct {
	*ordCards

	gaps []Cards

	groupKind []Cards
	groupSuit []Cards

	paired *map[int][]Cards
	suited *map[int][]Cards
}

type Hand struct {
	*handCards

	Rank   hand.Rank
	Value  Cards
	High   Cards
	Kicker Cards

	rank   bool
	high   bool
	kicker bool
}

func NewHandCards(o *ordCards) *handCards {
	groupKind := o.GroupByKind()
	groupSuit := o.GroupBySuit()

	return &handCards{
		ordCards: o,

		gaps: *o.Gaps(),

		groupKind: *groupKind,
		paired:    CountGroups(groupKind),

		groupSuit: *groupSuit,
		suited:    CountGroups(groupSuit),
	}
}

func (c *handCards) Cards() *Cards {
	return c.ordCards.Cards
}

func (c *handCards) Ordering() Ordering {
	return c.ordCards.Ordering
}

type rankFunc func(*handCards) (hand.Rank, *Hand)

func (c *handCards) Detect(ranks []rankFunc) *Hand {
	var result *Hand

	for _, r := range ranks {
		rank, hand := r(c)

		if hand != nil {
			if !hand.rank {
				hand.Rank = rank
			}
			if hand.high {
				hand.High = Cards{hand.Value[0]}
			}
			if hand.kicker {
				hand.Kicker = *c.ordCards.Kickers(&hand.Value)
			}

			hand.handCards = c

			result = hand

			break
		}
	}

	return result
}

func (h *Hand) String() string {
	return fmt.Sprintf("rank=%s high=%s value=%s kicker=%s",
		h.Rank,
		//h.hc.Cards(),
		h.High,
		h.Value,
		h.Kicker,
	)
}

func (h *Hand) ConsoleString() string {
	return fmt.Sprintf("rank=%s high=%s value=%s kicker=%s",
		h.Rank,
		//h.hc.Cards().ConsoleString(),
		h.High.ConsoleString(),
		h.Value.ConsoleString(),
		h.Kicker.ConsoleString(),
	)
}

type compareFunc func(*Hand, *Hand) int

var compareWith = func(ord Ordering) []compareFunc {
	return []compareFunc{
		func(a *Hand, b *Hand) int {
			return a.Rank.Compare(b.Rank)
		},

		func(a *Hand, b *Hand) int {
			return a.High.Compare(b.High, ord)
		},

		func(a *Hand, b *Hand) int {
			return a.Value.Compare(b.Value, ord)
		},

		func(a *Hand, b *Hand) int {
			return a.Kicker.Compare(b.Kicker, ord)
		},
	}
}

func (a *Hand) Compare(b *Hand) int {
	ord := a.handCards.Ordering()

	for _, compare := range compareWith(ord) {
		result := compare(a, b)
		if result != 0 {
			return result
		}
	}

	return 0
}
