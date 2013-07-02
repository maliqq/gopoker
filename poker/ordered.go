package poker

type Ordering int

type OrderedCards struct {
	value *Cards
	ord   Ordering
}

func (o *OrderedCards) Gaps() *[]Cards {
	sorted := make(Cards, len(*o.value))
	copy(sorted, *o.value)
	sort.Sort(ByKind{sorted, o.ord})

	cards := Cards{}
	for _, card := range *o.value {
		if Ace == card.kind {
			cards = append(cards, card)
		}
	}

	cards = append(cards, sorted...)
	return cards.GroupCards(func(card *Card, prev *Card) int {
		d := card.Index(o.ord) - prev.Index(o.ord)
		if d == 0 {
			return -1
		}
		if d == 1 {
			return 1
		}
		return 0
	})
}

func (o *OrderedCards) Kickers(cards *Cards) *Cards {
	length := 5 - len(*cards)

	diff := DiffCards(o.value, cards)
	sort.Sort(Arrange{ByKind{*diff, o.ord}})

	result := (*diff)[0:length]
	return &result
}

func (o *OrderedCards) GroupedByKind() *[]Cards {
	cards := make(Cards, len(*o.value))
	copy(cards, *o.value)
	sort.Sort(ByKind{cards, o.ord})

	return cards.GroupCards(func(card *Card, prev *Card) int {
		if card.kind == prev.kind {
			return 1
		}
		return 0
	})
}

func (o *OrderedCards) GroupedBySuit() *[]Cards {
	cards := make(Cards, len(*o.value))
	copy(cards, *o.value)
	sort.Sort(BySuit{cards})

	return cards.GroupCards(func(card *Card, prev *Card) int {
		if card.suit == prev.suit {
			return 1
		}
		return 0
	})
}
