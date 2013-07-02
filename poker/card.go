package poker

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

import (
	"gopoker/poker/card"
	"gopoker/util/console"
)

type Card struct {
	kind card.Kind
	suit card.Suit
}

type Cards []Card
type Ordering int

const (
	Ace              = card.Ace
	AceHigh Ordering = 0
	AceLow  Ordering = 1
)

type OrderedCards struct {
	value *Cards
	ord   Ordering
}

const (
	CardsNum = card.KindsNum * card.SuitsNum
)

func (card Card) String() string {
	return card.kind.String() + card.suit.String()
}

func (card Card) UnicodeString() string {
	return card.kind.String() + card.suit.UnicodeString()
}

func (c Card) Int() int {
	return (int(c.kind) << 2) + int(c.suit)
}

func (c Card) Byte() byte {
	return (byte(c.kind) << 2) + byte(c.suit)
}

func (c Card) ConsoleString() string {
	return fmt.Sprintf("%s%s%s", card.Colors[c.suit], c.UnicodeString(), console.RESET)
}

func (c Cards) String() string {
	s := ""
	for _, card := range c {
		s += card.String()
	}
	return s
}

func (c Cards) Binary() []byte {
	b := make([]byte, len(c))
	for i, card := range c {
		b[i] = card.Byte()
	}
	return b
}

func (c Cards) UnicodeString() string {
	s := ""
	for _, card := range c {
		s += card.UnicodeString()
	}
	return s
}

func (c Cards) ConsoleString() string {
	s := ""
	for _, card := range c {
		s += card.ConsoleString() + " "
	}
	return s
}

func (c Cards) BinaryString() string {
	return string(c.Binary())
}

func AllCards() *Cards {
	kinds := card.AllKinds()
	suits := card.AllSuits()
	cards := make(Cards, CardsNum)

	k := 0
	for _, kind := range kinds {
		for _, suit := range suits {
			cards[k] = Card{kind, suit}
			k++
		}
	}

	return &cards
}

func NewCard(i byte) (*Card, error) {
	if i < 0 || int(i) >= CardsNum {
		return nil, errors.New("invalid card")
	}
	return &Card{card.Kind(i >> 2), card.Suit(i % 4)}, nil
}

func MakeCard(kind int, suit int) (*Card, error) {
	k, err1 := card.MakeKind(kind)
	if err1 != nil {
		return nil, err1
	}
	s, err2 := card.MakeSuit(suit)
	if err2 != nil {
		return nil, err2
	}
	return &Card{k, s}, nil
}

func ParseCard(s string) (*Card, error) {
	if len(s) == 2 {
		p := strings.Split(s, "")

		kind := strings.Index(card.Kinds, strings.ToUpper(p[0]))
		suit := strings.Index(card.Suits, strings.ToLower(p[1]))

		card, err := MakeCard(kind, suit)
		if err != nil {
			return nil, err
		}

		return card, nil
	}
	return nil, errors.New(fmt.Sprintf("can't parse card %s", s))
}

func ParseCards(s string) (*Cards, error) {
	r, _ := regexp.Compile("(?i)[akqjt2-9]{1}[schd]{1}")
	match := r.FindAllString(s, len(s)/2+1)
	cards := make(Cards, len(match))
	i := 0
	for _, c := range match {
		card, err := ParseCard(c)
		if err != nil {
			return nil, err
		}
		cards[i] = *card
		i++
	}
	return &cards, nil
}

func ParseBinary(s []byte) (*Cards, error) {
	cards := make(Cards, len(s))
	for i, b := range s {
		card, err := NewCard(b)
		if err != nil {
			return nil, err
		}
		cards[i] = *card
	}
	return &cards, nil
}

func (c Cards) MarshalJSON() ([]byte, error) {
	//if len(c) == 0 {
	//  return []byte("null"), nil
	//}
	return []byte(strconv.Quote(c.BinaryString())), nil
}

func (c Card) Index(ord Ordering) int {
	switch ord {
	case AceHigh:
		return int(c.kind) + 1
	case AceLow:
		if Ace == c.kind {
			return 0
		}
		return int(c.kind) + 1
	}
	return -1
}

func (c Card) Equal(other Card) bool {
	return (c.kind == other.kind) && (c.suit == other.suit)
}

func ShuffleCards(c *Cards) *Cards {
	// seed random
	rand.Seed(time.Now().UnixNano())

	cards := *c
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return &cards
}

func GenerateCards(n int) *Cards {
	deck := NewDeck()
	cards := (*deck)[0:n]
	return &cards
}

func NewDeck() *Cards {
	return ShuffleCards(AllCards())
}

func DiffCards(a *Cards, b *Cards) *Cards {
	result := make(Cards, len(*a))
	present := make(map[int]bool, len(*b))

	for _, card := range *b {
		present[card.Int()] = true
	}

	i := 0
	for _, card := range *a {
		if _, err := present[card.Int()]; err == false {
			result[i] = card
			i++
		}
	}

	slice := result[0:i]

	return &slice
}

type groupFunc func(card *Card, prev *Card) int

func (cards *Cards) GroupCards(test groupFunc) *[]Cards {
	length := len(*cards)
	groups := make([]Cards, length)
	group := make(Cards, length)

	j, k := 0, 0
	for i := 0; i < length; i++ {
		card := (*cards)[i]

		if i == 0 {
			group[j] = card
			j++
		} else {
			prev := (*cards)[i-1]
			result := test(&card, &prev)

			if result == 1 {
				group[j] = card
				j++
			} else if result == 0 {
				groups[k] = group[0:j]
				k++
				group = make(Cards, length-i)
				group[0] = card
				j = 1
			}
		}
	}

	if j > 0 {
		groups[k] = group[0:j]
		k++
	}

	result := groups[0:k]

	return &result
}

func (c Cards) CombinePairs() []Cards {
	result := make([]Cards, len(c)/2)

	k := 0
	for i, first := range c {
		for j := i + 1; j < len(c); j++ {
			second := c[j]
			result[k] = Cards{first, second}
			k++
		}
	}

	return result
}

func CountGroups(groups *[]Cards) *map[int][]Cards {
	count := map[int][]Cards{}
	for _, group := range *groups {
		length := len(group)
		_, present := count[length]
		if !present {
			count[length] = []Cards{}
		}
		count[length] = append(count[length], group)
	}
	return &count
}

//
// sort
//
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

//
// OrderedCards
//
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
