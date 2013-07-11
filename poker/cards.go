package poker

import (
	"encoding/json"
	_ "fmt"
	"math/rand"
	"regexp"
	"time"
)

import (
	"gopoker/poker/card"
	"gopoker/util"
)

type Cards []Card

func AllCards() *Cards {
	all := card.All
	cards := make(Cards, len(all))
	for i, tuple := range all {
		cards[i] = Card{tuple.Kind, tuple.Suit}
	}
	return &cards
}

func GenerateCards(n int) *Cards {
	deck := NewDeck()
	cards := (*deck)[0:n]
	return &cards
}

func NewDeck() *Cards {
	return AllCards().Shuffle()
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

func (c Cards) MarshalJSON() ([]byte, error) {
	//if len(c) == 0 {
	//  return []byte("null"), nil
	//}
	return json.Marshal(c.UnicodeString())
}

func (this *Cards) Shuffle() *Cards {
	// seed random
	rand.Seed(time.Now().UnixNano())

	cards := *this
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return &cards
}

func (a *Cards) Diff(b *Cards) *Cards {
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

func (this *Cards) Group(test groupFunc) *[]Cards {
	length := len(*this)
	groups := make([]Cards, length)
	group := make(Cards, length)

	j, k := 0, 0
	for i := 0; i < length; i++ {
		card := (*this)[i]

		if i == 0 {
			group[j] = card
			j++

		} else {
			prev := (*this)[i-1]
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

func (this *Cards) Combine(m int) []Cards {
	n := len(*this)

	indexes := util.Comb(n, m)

	result := make([]Cards, len(indexes))

	for i, index := range indexes {
		cards := make(Cards, m)
		for i, j := range index {
			cards[i] = (*this)[j]
		}
		result[i] = cards
	}

	return result
}

func CountGroups(groups *[]Cards) *map[int][]Cards {
	count := map[int][]Cards{}

	for _, group := range *groups {
		length := len(group)
		if _, present := count[length]; !present {
			count[length] = []Cards{}
		}

		count[length] = append(count[length], group)
	}

	return &count
}
