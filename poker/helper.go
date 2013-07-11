package poker

import (
	"sort"
)

import (
	"gopoker/poker/card"
)

type Ordering int

const (
	Ace              = card.Ace
	AceHigh Ordering = 0
	AceLow  Ordering = 1
)

// cards with ordering (ace high/ace low)
type cardsHelper struct {
	Cards
	Ordering
	Reversed bool
}

func (this *cardsHelper) Qualify(q card.Kind) *Cards {
	qualified := Cards{}

	for _, card := range this.Cards {
		if card.Index(this.Ordering) < kindIndex(q, this.Ordering) {
			qualified = append(qualified, card)
		}
	}

	return &qualified
}

func (this *cardsHelper) Gaps() *GroupedCards {
	sorted := this.Arrange()

	cards := Cards{}
	for _, card := range this.Cards {
		if Ace == card.kind {
			cards = append(cards, card)
		}
	}

	cards = append(cards, *sorted...)

	return cards.Group(func(card *Card, prev *Card) int {
		d := card.Index(this.Ordering) - prev.Index(this.Ordering)

		if d == 0 {
			return -1
		}

		if d == 1 {
			return 1
		}

		return 0
	})
}

func (this *cardsHelper) Kickers(cards *Cards) *Cards {
	length := 5 - len(*cards)

	diff := this.Cards.Diff(cards)

	result := diff.Arrange(this.Ordering)
	result = result[0:length]

	return &result
}

func (this *cardsHelper) GroupByKind() *GroupedCards {
	sorted := this.Cards.Arrange(this.Ordering)

	return sorted.Group(func(card *Card, prev *Card) int {
		if card.kind == prev.kind {
			return 1
		}

		return 0
	})
}

func (this *cardsHelper) GroupBySuit() *GroupedCards {
	cards := make(Cards, len(this.Cards))

	copy(cards, this.Cards)

	sort.Sort(BySuit{cards})

	return cards.Group(func(card *Card, prev *Card) int {
		if card.suit == prev.suit {
			return 1
		}

		return 0
	})
}

func (this *cardsHelper) Arrange() *Cards {
	cards := this.Cards.Arrange(this.Ordering)

	return &cards
}

func (this *cardsHelper) Reverse() *Cards {
	cards := this.Cards.Reverse(this.Ordering)

	return &cards
}
