package poker

import (
	"fmt"
	"strings"
)

import (
	"gopoker/exch/message"
	"gopoker/poker/hand"
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

// Hand - poker hand
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

// NewHandCards - create indexed hand cards
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

// String - hand cards to string
func (hc *handCards) String() string {
	return fmt.Sprintf("gaps=%s paired=%s suited=%s", hc.gaps, hc.paired, hc.suited)
}

type rankFunc func(*handCards) (hand.Rank, *Hand)

// Detect - detect rank of hand cards
func (hc *handCards) Detect(ranks []rankFunc) *Hand {
	var result *Hand

	for _, r := range ranks {
		rank, hand := r(hc)

		if hand != nil {
			if !hand.rank {
				hand.Rank = rank
			}
			if hand.high {
				hand.High = Cards{hand.Value[0]}
			}
			if hand.kicker {
				hand.Kicker = hc.cardsHelper.Kickers(hand.Value)
			}

			hand.handCards = hc

			result = hand

			break
		}
	}

	return result
}

// RankName - rank name for hand, e.g. "StraghtFlush"
func (h *Hand) RankName() string {
	return string(h.Rank)
}

// RankTitle - rank title
func (h *Hand) RankTitle() string {
	return strings.Title(h.RankName())
}

// String - hand to string
func (h *Hand) String() string {
	return fmt.Sprintf("rank=%s high=%s value=%s kicker=%s",
		h.Rank,
		//h.hc.Cards(),
		h.High,
		h.Value,
		h.Kicker,
	)
}

// ConsoleString - hand to console string
func (h *Hand) ConsoleString() string {
	return fmt.Sprintf("rank=%s high=%s value=%s kicker=%s",
		h.Rank,
		//h.hc.Cards().ConsoleString(),
		h.High.ConsoleString(),
		h.Value.ConsoleString(),
		h.Kicker.ConsoleString(),
	)
}

// PrintString - hand to human readable string
func (h *Hand) PrintString() string {
	switch h.Rank {
	case hand.HighCard:
		return fmt.Sprintf("high card %s",
			h.High[0].KindTitle(),
		)

	case hand.OnePair:
		return fmt.Sprintf("pair of %ss",
			h.High[0].KindTitle(),
		)

	case hand.TwoPair:
		return fmt.Sprintf("two pairs, %ss and %ss",
			h.High[0].KindTitle(),
			h.High[1].KindTitle(),
		)

	case hand.ThreeKind:
		return fmt.Sprintf("three of a kind, %ss",
			h.High[0].KindTitle(),
		)

	case hand.Straight:
		return fmt.Sprintf("straight, %s to %s",
			h.Value.Min(AceHigh).KindTitle(),
			h.Value.Max(AceHigh).KindTitle(),
		)

	case hand.Flush:
		return fmt.Sprintf("flush, %s high",
			h.High[0].KindTitle(),
		)

	case hand.FullHouse:
		return fmt.Sprintf("full house, %ss full of %ss",
			h.High[0].KindTitle(),
			h.High[1].KindTitle(),
		)

	case hand.FourKind:
		return fmt.Sprintf("four of a kind, %ss",
			h.High[0].KindTitle(),
		)

	case hand.StraightFlush:
		return fmt.Sprintf("straight flush, %s to %s",
			h.Value.Min(AceHigh).KindTitle(),
			h.Value.Max(AceHigh).KindTitle(),
		)

	case hand.BadugiOne:
		return fmt.Sprintf("1-card badugi: %s",
			h.Value[0].KindTitle(),
		)

	case hand.BadugiTwo:
		return fmt.Sprintf("2-card badugi: %s + %s",
			h.Value[0].KindTitle(),
			h.Value[1].KindTitle(),
		)

	case hand.BadugiThree:
		return fmt.Sprintf("3-card badugi: %s + %s + %s",
			h.Value[0].KindTitle(),
			h.Value[1].KindTitle(),
			h.Value[2].KindTitle(),
		)

	case hand.BadugiFour:
		return fmt.Sprintf("4-card badugi: %s + %s + %s + %s",
			h.Value[0].KindTitle(),
			h.Value[1].KindTitle(),
			h.Value[2].KindTitle(),
			h.Value[3].KindTitle(),
		)
	}

	return ""
}

// Proto - protobuf representation of hand
func (h *Hand) Proto() *message.Hand {
	return &message.Hand{
		Rank: message.Rank(
			message.Rank_value[string(h.Rank)],
		).Enum(),

		High:   h.High.Binary(),
		Value:  h.Value.Binary(),
		Kicker: h.Kicker.Binary(),
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

// Compare - compare two hands by rank, high cards, value, kickers
func (h *Hand) Compare(o *Hand) int {
	ord := h.handCards.Ordering

	for _, compare := range compareWith(ord) {
		result := compare(h, o)
		if result != 0 {
			return result
		}
	}

	return 0
}
