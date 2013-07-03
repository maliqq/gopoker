package poker

import (
	"fmt"
	"gopoker/poker/hand"
)

type PocketCards struct {
	cards *OrderedCards

	gaps []Cards

	groupKind []Cards
	groupSuit []Cards

	paired *map[int][]Cards
	suited *map[int][]Cards
}

type Hand struct {
	pocket *PocketCards

	Rank   hand.Rank
	Value  Cards
	High   Cards
	Kicker Cards

	rank   bool
	high   bool
	kicker bool
}

// Pocket
func NewPocket(o *OrderedCards) *PocketCards {
	groupKind := o.GroupedByKind()
	groupSuit := o.GroupedBySuit()

	return &PocketCards{
		cards: o,

		gaps: *o.Gaps(),

		groupKind: *groupKind,
		paired:    CountGroups(groupKind),

		groupSuit: *groupSuit,
		suited:    CountGroups(groupSuit),
	}
}

func (p *PocketCards) Cards() *Cards {
	return p.cards.value
}

func (p *PocketCards) Ordering() Ordering {
	return p.cards.ord
}

type rankFunc func(*PocketCards)(hand.Rank, *Hand)

func (pocket *PocketCards) Detect(ranks []rankFunc) *Hand {
	var result *Hand

	for _, r := range ranks {
		rank, hand := r(pocket)

		if hand != nil {
			if !hand.rank {
				hand.Rank = rank
			}
			if hand.high {
				hand.High = Cards{hand.Value[0]}
			}
			if hand.kicker {
				hand.Kicker = *pocket.cards.Kickers(&hand.Value)
			}

			hand.pocket = pocket

			result = hand

			break
		}
	}

	return result
}

func (h *Hand) String() string {
	return fmt.Sprintf("rank=%s high=%s value=%s kicker=%s",
		h.Rank,
		//h.pocket.Cards(),
		h.High,
		h.Value,
		h.Kicker,
	)
}

func (h *Hand) ConsoleString() string {
	return fmt.Sprintf("rank=%s high=%s value=%s kicker=%s",
		h.Rank,
		//h.pocket.Cards().ConsoleString(),
		h.High.ConsoleString(),
		h.Value.ConsoleString(),
		h.Kicker.ConsoleString(),
	)
}

type compareFunc func(*Hand, *Hand) int

func (a *Hand) Compare(b *Hand) int {
	ord := a.pocket.Ordering()

	comparers := []compareFunc{
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

	for _, f := range comparers {
		result := f(a, b)
		if result != 0 {
			return result
		}
	}

	return 0
}

type ByHand struct {
	hands []*Hand
}

func (h ByHand) Len() int {
	return len(h.hands)
}

func (h ByHand) Swap(i, j int) {
	h.hands[i], h.hands[j] = h.hands[j], h.hands[i]
}

func (h ByHand) Less(i, j int) bool {
	a := h.hands[i]
	b := h.hands[j]

	return a.Compare(b) == -1
}
