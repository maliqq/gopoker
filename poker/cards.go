package poker

import (
	_ "bytes"
	"encoding/json"
	"math/rand"
	"regexp"
	"sort"
	"time"
)

import (
	"gopoker/poker/card"
	"gopoker/protocol/message"
	"gopoker/util"
)

type Cards []*Card

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

var All = AllCards()

func GenerateCards(n int) Cards {
	deck := NewDeck()

	return deck[0:n]
}

func NewDeck() Cards {
	return All.Shuffle()
}

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

func StringCards(s string) Cards {
	cards, _ := ParseCards(s)
	return cards
}

func BinaryCards(s []byte) Cards {
	cards, _ := ParseBinary(s)
	return cards
}

func (c Cards) Uint64() uint64 {
	var result uint64 = uint64(c[0].Index(AceHigh))
	for i := 1; i < len(c); i++ {
		result |= uint64(card.Masks[c[i].Index(AceHigh)])
	}
	return result
}

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

func (c Cards) Proto() message.Cards {
	return c.Binary()
}

func (c Cards) String() string {
	s := ""
	for _, card := range c {
		s += card.String()
	}

	return s
}

func (c Cards) PrintString() string {
	return c.String()
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

func (c Cards) MarshalJSON() ([]byte, error) {
	//if len(c) == 0 {
	//  return []byte("null"), nil
	//}

	//var buf bytes.Buffer
	//buf.Write(c.Binary())
	//return json.Marshal(buf.String())
	return json.Marshal(c.Binary())
}

func (this Cards) Shuffle() Cards {
	// seed random
	rand.Seed(time.Now().UnixNano())

	cards := this
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}

func (a Cards) Diff(b Cards) Cards {
	result := make(Cards, len(a))
	present := make(map[int]bool, len(b))

	for _, card := range b {
		present[card.Int()] = true
	}

	i := 0
	for _, card := range a {
		if _, err := present[card.Int()]; err == false {
			result[i] = card
			i++
		}
	}

	return result[0:i]
}

func (a Cards) Append(b Cards) Cards {
	c := make(Cards, len(a))
	copy(c, a)
	return append(c, b...)
}

type groupFunc func(card *Card, prev *Card) int

func (this Cards) Group(test groupFunc) GroupedCards {
	length := len(this)
	groups := make(GroupedCards, length)
	group := make(Cards, length)

	j, k := 0, 0
	for i := 0; i < length; i++ {
		card := this[i]

		if i == 0 {
			group[j] = card
			j++

		} else {
			prev := this[i-1]
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

func (this Cards) Combine(m int) GroupedCards {
	n := len(this)

	indexes := util.Combine(n, m)

	result := make(GroupedCards, len(indexes))

	for i, index := range indexes {
		cards := make(Cards, m)
		for i, j := range index {
			cards[i] = this[j]
		}
		result[i] = cards
	}

	return result
}

func (a Cards) Equal(b Cards) bool {
	if len(a) != len(b) {
		return false
	}

	for i, card := range a {
		if !card.Equal(a[i]) {
			return false
		}
	}

	return true
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
}

func (c Cards) Arrange(ord Ordering) Cards {
	sort.Sort(Arrange{ByKind{c, ord}})

	return c
}

func (c Cards) Reverse(ord Ordering) Cards {
	sort.Sort(Reverse{ByKind{c, ord}})

	return c
}

type maxFunc func(d int) bool

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

func (c Cards) Min(ord Ordering) *Card {
	return c.MaxBy(ord, func(d int) bool {
		return d < 0
	})
}

func (c Cards) Max(ord Ordering) *Card {
	return c.MaxBy(ord, func(d int) bool {
		return d > 0
	})
}

func (c Cards) IsPair() bool {
	return len(c) == 2 && c[0].kind == c[1].kind
}

func (c Cards) IsSuited() bool {
	return len(c) == 2 && c[0].suit == c[1].suit
}

//func (p *Pocket) IsConnector() bool {
//
//}
