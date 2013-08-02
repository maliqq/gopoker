package poker

// sort cards

// Len - len
func (c Cards) Len() int {
	return len(c)
}

// Swap - swap
func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// BySuit - sort by suit (suit sensitive sort)
type BySuit struct {
	Cards
}

// Less - less
func (c BySuit) Less(i, j int) bool {
	return c.Cards[i].suit < c.Cards[j].suit
}

// ByKind - sort by kind
type ByKind struct {
	Cards
	Ordering
}

// Len - len
func (c ByKind) Len() int {
	return len(c.Cards)
}

// Swap - swap
func (c ByKind) Swap(i, j int) {
	c.Cards.Swap(i, j)
}

// Less - less
func (c ByKind) Less(i, j int) bool {
	card1 := c.Cards[i]
	card2 := c.Cards[j]

	return card1.Compare(card2, c.Ordering) == -1
}

// ByFirst - sort by first in group
type ByFirst struct {
	GroupedCards
	Ordering
}

// Len - len
func (c ByFirst) Len() int {
	return len(c.GroupedCards)
}

// Swap - swap
func (c ByFirst) Swap(i, j int) {
	c.GroupedCards[i], c.GroupedCards[j] = c.GroupedCards[j], c.GroupedCards[i]
}

// Less - less
func (c ByFirst) Less(i, j int) bool {
	card1 := c.GroupedCards[i][0]
	card2 := c.GroupedCards[j][0]

	return card2.Compare(card1, c.Ordering) == -1
}

// ByMax - sort by max in group
type ByMax struct {
	GroupedCards
	Ordering
}

// Len - len
func (c ByMax) Len() int {
	return len(c.GroupedCards)
}

// Swap - swap
func (c ByMax) Swap(i, j int) {
	c.GroupedCards[i], c.GroupedCards[j] = c.GroupedCards[j], c.GroupedCards[i]
}

// Less - less
func (c ByMax) Less(i, j int) bool {
	max1 := c.GroupedCards[i].Max(c.Ordering)
	max2 := c.GroupedCards[j].Max(c.Ordering)

	return max2.Compare(max1, c.Ordering) == -1
}

//
// sort hand
//

// ByHand - sort by hand
type ByHand struct {
	hands []*Hand
}

// Len - len
func (h ByHand) Len() int {
	return len(h.hands)
}

// Swap - swap
func (h ByHand) Swap(i, j int) {
	h.hands[i], h.hands[j] = h.hands[j], h.hands[i]
}

// Less - less
func (h ByHand) Less(i, j int) bool {
	a := h.hands[i]
	b := h.hands[j]

	return a.Compare(b) == -1
}

// Reverse - arranging cards in reverse order
type Reverse struct {
	ByKind
}

// Less - less
func (c Reverse) Less(i, j int) bool {
	return c.ByKind.Less(i, j)
}

// Arrange - arranging cards - direct order
type Arrange struct{ ByKind }

// Less -l ess
func (c Arrange) Less(i, j int) bool {
	return c.ByKind.Less(j, i)
}
