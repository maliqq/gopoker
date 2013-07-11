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

func GenerateCards(n int) *Cards {
	deck := NewDeck()
	cards := (*deck)[0:n]
	return &cards
}

func NewDeck() *Cards {
	return ShuffleCards(AllCards())
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
	return json.Marshal(c.Binary())
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

func (c Cards) Combine(m int) []Cards {
	n := len(c)
	size := util.Fact(int64(n)) / util.Fact(int64(n-m)) / util.Fact(int64(m))
	result := make([]Cards, size)

	index := make([]int, m)
	for i := range index {
		index[i] = i
	}

	k := 0
	cards := make(Cards, m)
	for i, j := range index[:m] {
		cards[i] = c[j]
	}
	result[k] = cards
	k++

	for {
		i := m - 1
		for ; i >= 0 && index[i] == i+n-m; i -= 1 {
		}

		if i < 0 {
			break
		}

		index[i] += 1
		for j := i + 1; j < m; j += 1 {
			index[j] = index[j-1] + 1
		}

		cards := make(Cards, m)
		for i, j := range index {
			cards[i] = c[j]
		}
		result[k] = cards
		k++
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
