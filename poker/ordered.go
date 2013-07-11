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
type ordCards struct {
	*Cards
	Ordering
}

func NewOrderedCards(cards *Cards, ord Ordering) *ordCards {
	return &ordCards{
		Cards: cards,
		Ordering: ord,
	}
}

func (this *ordCards) Qualify(q card.Kind) *ordCards {
	qualified := Cards{}

	for _, card := range *this.Cards {
		if card.Index(this.Ordering) < kindIndex(q, this.Ordering) {
			qualified = append(qualified, card)
		}
	}

	return NewOrderedCards(&qualified, this.Ordering)
}

func (this *ordCards) Gaps() *[]Cards {
	sorted := make(Cards, len(*this.Cards))

	copy(sorted, *this.Cards)

	sort.Sort(ByKind{sorted, this.Ordering})

	cards := Cards{}
	for _, card := range *this.Cards {
		if Ace == card.kind {
			cards = append(cards, card)
		}
	}

	cards = append(cards, sorted...)

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

func (this *ordCards) Kickers(cards *Cards) *Cards {
	length := 5 - len(*cards)

	diff := this.Cards.Diff(cards)

	sort.Sort(Arrange{ByKind{*diff, this.Ordering}})

	result := (*diff)[0:length]

	return &result
}

func (this *ordCards) GroupByKind() *[]Cards {
	cards := make(Cards, len(*this.Cards))

	copy(cards, *this.Cards)

	sort.Sort(ByKind{cards, this.Ordering})

	return cards.Group(func(card *Card, prev *Card) int {
		if card.kind == prev.kind {
			return 1
		}

		return 0
	})
}

func (this *ordCards) GroupBySuit() *[]Cards {
	cards := make(Cards, len(*this.Cards))

	copy(cards, *this.Cards)

	sort.Sort(BySuit{cards})

	return cards.Group(func(card *Card, prev *Card) int {
		if card.suit == prev.suit {
			return 1
		}

		return 0
	})
}
