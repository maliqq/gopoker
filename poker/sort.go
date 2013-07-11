package poker

// sort cards
func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// sort by suit (suit sensitive sort)
type BySuit struct {
	Cards
}

func (c BySuit) Less(i, j int) bool {
	return c.Cards[i].suit < c.Cards[j].suit
}

// sort by kind
func (c ordCards) Len() int {
	return len(*c.Cards)
}

func (c ordCards) Swap(i, j int) {
	c.Cards.Swap(i, j)
}

func (c ordCards) Less(i, j int) bool {
	card1 := (*c.Cards)[i]
	card2 := (*c.Cards)[j]

	return card1.Compare(card2, c.Ordering) == -1
}

// sort by first in group
type ByFirst struct {
	GroupedCards
	Ordering
}

func (c ByFirst) Len() int {
	return len(c.GroupedCards)
}

func (c ByFirst) Swap(i, j int) {
	c.GroupedCards[i], c.GroupedCards[j] = c.GroupedCards[j], c.GroupedCards[i]
}

func (c ByFirst) Less(i, j int) bool {
	card1 := c.GroupedCards[i][0]
	card2 := c.GroupedCards[j][0]

	return card2.Compare(card1, c.Ordering) == -1
}

// sort by max in group
type ByMax struct {
	GroupedCards
	Ordering
}

func (c ByMax) Len() int {
	return len(c.GroupedCards)
}

func (c ByMax) Swap(i, j int) {
	c.GroupedCards[i], c.GroupedCards[j] = c.GroupedCards[j], c.GroupedCards[i]
}

func (c ByMax) Less(i, j int) bool {
	max1 := c.GroupedCards[i].Max(c.Ordering)
	max2 := c.GroupedCards[j].Max(c.Ordering)

	return max2.Compare(*max1, c.Ordering) == -1
}

//
// sort hand
//
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

// arranging cards - reverse order
type Reverse struct {
	ordCards
}

func (c Reverse) Less(i, j int) bool {
	return c.ordCards.Less(i, j)
}

// arranging cards - direct order
type Arrange struct{ ordCards }

func (c Arrange) Less(i, j int) bool {
	return c.ordCards.Less(j, i)
}
