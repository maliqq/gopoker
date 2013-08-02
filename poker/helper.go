package poker

import (
	"sort"
)

import (
	"gopoker/poker/card"
	"gopoker/poker/hand"
)

// Ordering - card ordering
type Ordering int

// Orderings
const (
	Ace              = card.Ace // Ace - kind ace
	AceHigh Ordering = 0        // AceHigh - ace is high card
	AceLow  Ordering = 1        // AceLow - ace is low card
)

// cards with ordering (ace high/ace low)
type cardsHelper struct {
	Cards
	Ordering
	Low bool
}

// Qualify - filter cards by qualifier card
func (helper *cardsHelper) Qualify(q card.Kind) {
	qualified := Cards{}

	for _, card := range helper.Cards {
		if card.Index(helper.Ordering) <= kindIndex(q, helper.Ordering) {
			qualified = append(qualified, card)
		}
	}

	helper.Cards = qualified
}

// Gaps - group cards by gaps
func (helper *cardsHelper) Gaps() GroupedCards {
	sorted := helper.Reverse()

	cards := Cards{}
	for _, card := range helper.Cards {
		if Ace == card.kind {
			cards = append(cards, card)
		}
	}

	cards = append(cards, sorted...)

	return cards.Group(func(card *Card, prev *Card) int {
		d := card.Index(helper.Ordering) - prev.Index(helper.Ordering)

		if d == 0 {
			return -1
		}

		if d == 1 || d == -12 {
			return 1
		}

		return 0
	})
}

// Kickers - get kickers
func (helper *cardsHelper) Kickers(cards Cards) Cards {
	length := 5 - len(cards)

	diff := helper.Cards.Diff(cards)

	result := diff.Arrange(helper.Ordering)
	result = result[0:length]

	return result
}

// GroupByKind - group by card kinds
func (helper *cardsHelper) GroupByKind() GroupedCards {
	sorted := helper.Cards.Arrange(helper.Ordering)

	return sorted.Group(func(card *Card, prev *Card) int {
		if card.kind == prev.kind {
			return 1
		}

		return 0
	})
}

// GroupBySuit - group by card suits
func (helper *cardsHelper) GroupBySuit() GroupedCards {
	cards := make(Cards, len(helper.Cards))

	copy(cards, helper.Cards)

	sort.Sort(BySuit{cards})

	return cards.Group(func(card *Card, prev *Card) int {
		if card.suit == prev.suit {
			return 1
		}

		return 0
	})
}

// Arrange - sort cards by ordering
func (helper *cardsHelper) Arrange() Cards {
	return helper.Cards.Arrange(helper.Ordering)
}

// Reverse - reverse sort cards by ordering
func (helper *cardsHelper) Reverse() Cards {
	return helper.Cards.Reverse(helper.Ordering)
}

// IsLow - check if hand is low
func (helper *cardsHelper) IsLow() (*Hand, error) {
	uniq := Cards{}
	for _, cards := range helper.GroupByKind() {
		uniq = append(uniq, cards[0])
	}

	lowCards := uniq.Reverse(helper.Ordering)

	if len(lowCards) == 0 {
		return nil, nil
	}

	if len(lowCards) >= 5 {
		lowCards = lowCards[0:5]
	}

	max := lowCards.Max(helper.Ordering)
	newHand := &Hand{
		Value: lowCards,
		High:  Cards{max},
	}

	if len(lowCards) == 5 {
		newHand.Rank = hand.CompleteLow
	} else {
		newHand.Rank = hand.IncompleteLow
	}

	return newHand, nil
}

// IsGapLow - check if hand is gapped low
func (helper *cardsHelper) IsGapLow() (*Hand, error) {
	high, err := isHigh(&helper.Cards)
	if err != nil {
		return nil, err
	}

	if high.Rank == hand.HighCard {
		return helper.IsLow()
	}

	return high, nil
}
