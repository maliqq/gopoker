package poker

type BySuit struct{ Cards }

type ByKind struct {
	cards Cards
	ord   Ordering
}

type ByFirst struct {
	groups []Cards
	ord    Ordering
}

type ByMax struct {
	groups []Cards
	ord    Ordering
}

type Arrange struct {
	ByKind
}

func (c Cards) Len() int {
	return len(c)
}

func (c ByFirst) Len() int {
	return len(c.groups)
}

func (c ByMax) Len() int {
	return len(c.groups)
}

type maxFunc func(d int) bool

func (card1 Card) Compare(card2 Card, ord Ordering) int {
	a, b := card1.Index(ord), card2.Index(ord)
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}

func (a Cards) Compare(b Cards, ord Ordering) int {
	if len(a) == len(b) {
		for i, left := range a {
			right := b[i]
			result := left.Compare(right, ord)
			if result != 0 {
				return result
			}
		}
		return 0
	} else {
		min := len(a)
		if len(b) < min {
			min = len(b)
		}
		return a[0:min].Compare(b[0:min], ord)
	}
	return 1
}

func (c Cards) MaxBy(ord Ordering, f maxFunc) *Card {
	result := &c[0]
	max := result.Index(ord)
	for _, card := range c {
		i := card.Index(ord)
		if f(i - max) {
			max = i
			result = &card
		}
	}
	return result
}

func (c Cards) Max(ord Ordering) *Card {
	return c.MaxBy(ord, func(d int) bool {
		return d > 0
	})
}

func (c Cards) Min(ord Ordering) *Card {
	return c.MaxBy(ord, func(d int) bool {
		return d < 0
	})
}

func (c ByKind) Len() int {
	return len(c.cards)
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ByKind) Swap(i, j int) {
	c.cards.Swap(i, j)
}

func (c ByFirst) Swap(i, j int) {
	c.groups[i], c.groups[j] = c.groups[j], c.groups[i]
}

func (c ByMax) Swap(i, j int) {
	c.groups[i], c.groups[j] = c.groups[j], c.groups[i]
}

func (c BySuit) Less(i, j int) bool {
	return c.Cards[i].suit < c.Cards[j].suit
}

func (c ByKind) Less(i, j int) bool {
	card1 := c.cards[i]
	card2 := c.cards[j]
	return card1.Compare(card2, c.ord) == -1
}

func (c ByFirst) Less(i, j int) bool {
	card1 := c.groups[i][0]
	card2 := c.groups[j][0]
	return card2.Compare(card1, c.ord) == -1
}

func (c ByMax) Less(i, j int) bool {
	max1 := c.groups[i].Max(c.ord)
	max2 := c.groups[j].Max(c.ord)
	return max2.Compare(*max1, c.ord) == -1
}

func (c Arrange) Less(i, j int) bool {
	return c.ByKind.Less(j, i)
}

func ArrangeCards(c *Cards, ord Ordering) *Cards {
	cards := *c
	sort.Sort(Arrange{ByKind{cards, ord}})
	return &cards
}

func ArrangeGroupsByFirst(c *[]Cards, ord Ordering) *[]Cards {
	groups := *c
	sort.Sort(ByFirst{groups, ord})
	return &groups
}

func ArrangeGroupsByMax(c *[]Cards, ord Ordering) *[]Cards {
	groups := *c
	sort.Sort(ByMax{groups, ord})
	return &groups
}
