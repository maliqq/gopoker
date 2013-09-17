package poker

import (
	"encoding/json"
	"math/rand"
	"regexp"
	"sort"
	"time"
)

import (
	"gopoker/event/message/format/protobuf"
	"gopoker/poker/card"
	"gopoker/util"
)

// Cards - set of cards
type Cards []*Card

// AllCards - 52 cards deck
func AllCards() Cards {
	cards := make(Cards, card.CardsNum)

	k := 0
	for _, kind := range card.AllKinds() {
		for _, suit := range card.AllSuits() {
			cards[k] = &Card{kind, suit}
			k++
		}
	}

	return cards
}

// All - 52 cards deck
var All = AllCards()

// GenerateCards - generate random deck of n cards
func GenerateCards(n int) Cards {
	deck := NewDeck()

	return deck[0:n]
}

// NewDeck - create random 52 card deck
func NewDeck() Cards {
	return All.Shuffle()
}

// ParseCards - parse cards from string (format: [AKQJT0-9]{1}[shdc]{1})
func ParseCards(s string) (Cards, error) {
	r, _ := regexp.Compile("(?i)[akqjt2-9]{1}[schd]{1}")
	match := r.FindAllString(s, len(s)/2+1)

	cards := make(Cards, len(match))

	i := 0
	for _, c := range match {
		card, err := ParseCard(c)
		if err != nil {
			return nil, err
		}
		cards[i] = card
		i++
	}

	return cards, nil
}

// ParseBinary - parse cards from binary
func ParseBinary(s []byte) (Cards, error) {
	cards := make(Cards, len(s))
	for i, b := range s {
		card, err := NewCard(b)
		if err != nil {
			return nil, err
		}
		cards[i] = card
	}

	return cards, nil
}

// StringCards - factory for cards from string
func StringCards(s string) Cards {
	cards, _ := ParseCards(s)
	return cards
}

// BinaryCards - factory for cards from binary
func BinaryCards(s []byte) Cards {
	cards, _ := ParseBinary(s)
	return cards
}

// Uint64 - uint64 representation of card
func (c Cards) Uint64() uint64 {
	var result = uint64(c[0].Index(AceHigh))
	for i := 1; i < len(c); i++ {
		result |= uint64(card.Masks[c[i].Index(AceHigh)])
	}
	return result
}

// Binary - binary(byte) representation of card
func (c Cards) Binary() []byte {
	b := make([]byte, len(c))
	for i, card := range c {
		if card == nil {
			b[i] = 0
		} else {
			b[i] = card.Byte()
		}
	}

	return b
}

// Proto - protobuf representation of card
func (c Cards) Proto() protobuf.Cards {
	return c.Binary()
}

// String - string representation of cards
func (c Cards) String() string {
	s := ""
	for _, card := range c {
		s += card.String()
	}

	return s
}

// PrintString - human readable string of cards
func (c Cards) PrintString() string {
	return c.String()
}

// UnicodeString - unicode string of cards
func (c Cards) UnicodeString() string {
	s := ""
	for _, card := range c {
		s += card.UnicodeString()
	}

	return s
}

// ConsoleString - colorified string of cards
func (c Cards) ConsoleString() string {
	s := ""
	for _, card := range c {
		s += card.ConsoleString() + " "
	}

	return s
}

// MarshalJSON - JSON representation of card
func (c Cards) MarshalJSON() ([]byte, error) {
	//if len(c) == 0 {
	//  return []byte("null"), nil
	//}

	//var buf bytes.Buffer
	//buf.Write(c.Binary())
	//return json.Marshal(buf.String())
	return json.Marshal(c.Binary())
}

// Shuffle - randomly shuffle cards
func (c Cards) Shuffle() Cards {
	// seed random
	rand.Seed(time.Now().UnixNano())

	cards := c
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}

// Diff - difference between two sets of cards
func (c Cards) Diff(o Cards) Cards {
	result := make(Cards, len(c))
	present := make(map[int]bool, len(o))

	for _, card := range o {
		present[card.Int()] = true
	}

	i := 0
	for _, card := range c {
		if _, err := present[card.Int()]; err == false {
			result[i] = card
			i++
		}
	}

	return result[0:i]
}

// Append - append sets of cards
func (c Cards) Append(o Cards) Cards {
	buf := make(Cards, len(c))
	copy(buf, c)
	return append(buf, o...)
}

type groupFunc func(card *Card, prev *Card) int

// Group - group cards by function
func (c Cards) Group(test groupFunc) GroupedCards {
	length := len(c)
	groups := make(GroupedCards, length)
	group := make(Cards, length)

	j, k := 0, 0
	for i := 0; i < length; i++ {
		card := c[i]

		if i == 0 {
			group[j] = card
			j++

		} else {
			prev := c[i-1]
			result := test(card, prev)

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

	return groups[0:k]
}

// Combine - create combinations of cards
func (c Cards) Combine(m int) GroupedCards {
	n := len(c)

	indexes := util.Combine(n, m)

	result := make(GroupedCards, len(indexes))

	for i, index := range indexes {
		cards := make(Cards, m)
		for i, j := range index {
			cards[i] = c[j]
		}
		result[i] = cards
	}

	return result
}

// Equal - check equality of cards
func (c Cards) Equal(o Cards) bool {
	if len(c) != len(o) {
		return false
	}

	for i, card := range c {
		if !card.Equal(o[i]) {
			return false
		}
	}

	return true
}

// Compare - compare cards by ordering
func (c Cards) Compare(o Cards, ord Ordering) int {
	if len(c) == len(o) {
		for i, left := range c {
			right := o[i]

			result := left.Compare(right, ord)
			if result != 0 {
				return result
			}
		}

		return 0
	}

	min := len(c)

	if len(o) < min {
		min = len(o)
	}

	return c[0:min].Compare(o[0:min], ord)
}

// Arrange - arrange cards in ascending order
func (c Cards) Arrange(ord Ordering) Cards {
	sort.Sort(Arrange{ByKind{c, ord}})

	return c
}

// Reverse - arrange cards in descending order
func (c Cards) Reverse(ord Ordering) Cards {
	sort.Sort(Reverse{ByKind{c, ord}})

	return c
}

type maxFunc func(d int) bool

// MaxBy - largest card by function
func (c Cards) MaxBy(ord Ordering, f maxFunc) *Card {
	result := c[0]

	max := result.Index(ord)

	for _, card := range c {
		i := card.Index(ord)
		if f(i - max) {
			max = i
			result = card
		}
	}

	return result
}

// Min - smallest card by ordering
func (c Cards) Min(ord Ordering) *Card {
	return c.MaxBy(ord, func(d int) bool {
		return d < 0
	})
}

// Max - largest card by ordering
func (c Cards) Max(ord Ordering) *Card {
	return c.MaxBy(ord, func(d int) bool {
		return d > 0
	})
}

// IsPair - two cards with same kind
func (c Cards) IsPair() bool {
	return len(c) == 2 && c[0].kind == c[1].kind
}

// IsSuited - two cards with same suit
func (c Cards) IsSuited() bool {
	return len(c) == 2 && c[0].suit == c[1].suit
}

//func (p *Pocket) IsConnector() bool {
//
//}
