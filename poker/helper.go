package poker

import (
	"sort"
)

import (
	"gopoker/poker/card"
	"gopoker/poker/hand"
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
	Low bool
}

func (this *cardsHelper) Qualify(q card.Kind) {
	qualified := Cards{}

	for _, card := range this.Cards {
		if card.Index(this.Ordering) <= kindIndex(q, this.Ordering) {
			qualified = append(qualified, card)
		}
	}

	this.Cards = qualified
}

func (this *cardsHelper) Gaps() GroupedCards {
	sorted := this.Reverse()

	cards := Cards{}
	for _, card := range this.Cards {
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

		if d == 1 || d == -12 {
			return 1
		}

		return 0
	})
}

func (this *cardsHelper) Kickers(cards Cards) Cards {
	length := 5 - len(cards)

	diff := this.Cards.Diff(cards)

	result := diff.Arrange(this.Ordering)
	result = result[0:length]

	return result
}

func (this *cardsHelper) GroupByKind() GroupedCards {
	sorted := this.Cards.Arrange(this.Ordering)

	return sorted.Group(func(card *Card, prev *Card) int {
		if card.kind == prev.kind {
			return 1
		}

		return 0
	})
}

func (this *cardsHelper) GroupBySuit() GroupedCards {
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

func (this *cardsHelper) Arrange() Cards {
	return this.Cards.Arrange(this.Ordering)
}

func (this *cardsHelper) Reverse() Cards {
	return this.Cards.Reverse(this.Ordering)
}

func (this *cardsHelper) IsLow() (*Hand, error) {
	uniq := Cards{}
	for _, cards := range this.GroupByKind() {
		uniq = append(uniq, cards[0])
	}

	lowCards := uniq.Reverse(this.Ordering)

	if len(lowCards) == 0 {
		return nil, nil
	}

	if len(lowCards) >= 5 {
		lowCards = lowCards[0:5]
	}

	max := lowCards.Max(this.Ordering)
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

func (this *cardsHelper) IsGapLow() (*Hand, error) {
	high, err := isHigh(&this.Cards)
	if err != nil {
		return nil, err
	}

	if high.Rank == hand.HighCard {
		return this.IsLow()
	}

	return high, nil
}
