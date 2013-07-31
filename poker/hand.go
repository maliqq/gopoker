package poker

import (
	"fmt"
	"strings"
)

import (
	"gopoker/poker/hand"
	"gopoker/protocol/message"
)

type handCards struct {
	cardsHelper

	gaps GroupedCards

	groupKind GroupedCards
	groupSuit GroupedCards

	// fixme pointers
	paired map[int]GroupedCards
	suited map[int]GroupedCards
}

type Hand struct {
	handCards *handCards

	Rank   hand.Rank
	Value  Cards
	High   Cards
	Kicker Cards

	rank   bool
	high   bool
	kicker bool
}

func NewHandCards(cards *Cards, ord Ordering, reversed bool) *handCards {
	helper := cardsHelper{*cards, ord, reversed}

	groupKind := helper.GroupByKind()
	groupSuit := helper.GroupBySuit()

	hc := handCards{
		cardsHelper: helper,

		gaps: helper.Gaps(),

		groupKind: groupKind,
		paired:    groupKind.Count(),

		groupSuit: groupSuit,
		suited:    groupSuit.Count(),
	}

	return &hc
}

func (hc *handCards) String() string {
	return fmt.Sprintf("gaps=%s paired=%s suited=%s", hc.gaps, hc.paired, hc.suited)
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
				hand.Kicker = c.cardsHelper.Kickers(hand.Value)
			}

			hand.handCards = c

			result = hand

			break
		}
	}

	return result
}

func (h *Hand) RankName() string {
	return string(h.Rank)
}

func (h *Hand) RankTitle() string {
	return strings.Title(h.RankName())
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

// human readable string
func (h *Hand) PrintString() string {
	switch h.Rank {
	case hand.HighCard:
		return fmt.Sprintf("high card %s", h.High[0].KindTitle())

	case hand.OnePair:
		return fmt.Sprintf("pair of %ss", h.High[0].KindTitle())

	case hand.TwoPair:
		return fmt.Sprintf("two pairs, %ss and %ss", h.High[0].KindTitle(), h.High[1].KindTitle())

	case hand.ThreeKind:
		return fmt.Sprintf("three of a kind, %ss", h.High[0].KindTitle())

	case hand.Straight:
		return fmt.Sprintf("straight, %s to %s", h.Value.Min(AceHigh).KindTitle(), h.Value.Max(AceHigh).KindTitle())

	case hand.Flush:
		return fmt.Sprintf("flush, %s high", h.High[0].KindTitle())

	case hand.FullHouse:
		return fmt.Sprintf("full house, %ss full of %ss", h.High[0].KindTitle(), h.High[1].KindTitle())

	case hand.FourKind:
		return fmt.Sprintf("four of a kind, %ss", h.High[0].KindTitle())

	case hand.StraightFlush:
		return fmt.Sprintf("straight flush, %s to %s", h.Value.Min(AceHigh).KindTitle(), h.Value.Max(AceHigh).KindTitle())
	}

	return ""
}

func (hand *Hand) Proto() *message.Hand {
	return &message.Hand{
		Rank: message.Rank(
			message.Rank_value[string(hand.Rank)],
		).Enum(),

		High:   hand.High.Binary(),
		Value:  hand.Value.Binary(),
		Kicker: hand.Kicker.Binary(),
	}
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
	ord := a.handCards.Ordering

	for _, compare := range compareWith(ord) {
		result := compare(a, b)
		if result != 0 {
			return result
		}
	}

	return 0
}
