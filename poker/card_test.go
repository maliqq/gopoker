package poker

import (
	"testing"
)

import (
	"gopoker/poker/card"
)

func TestCard(t *testing.T) {
	kinds := card.AllKinds()
	suits := card.AllSuits()
	cards := AllCards()

	s1 := ""
	for _, kind := range kinds {
		s1 += kind.String()
	}
	t.Logf("Kinds: %s", s1)

	s2 := ""
	for _, suit := range suits {
		s2 += suit.String()
	}
	t.Logf("Suits: %s", s2)

	t.Logf("Cards: %s", cards)

	if len(suits) != 4 {
		t.FailNow()
	}

	if len(kinds) != 13 {
		t.FailNow()
	}

	if len(cards) != 52 {
		t.FailNow()
	}
}

func TestParse(t *testing.T) {
	card1, _ := ParseCard("2s")
	_card1, _ := NewCard(0)
	t.Logf("card1=%s", *card1)
	if !_card1.Equal(*card1) {
		t.FailNow()
	}

	card2, _ := ParseCard("Ac")
	_card2, _ := NewCard(51)
	t.Logf("card2=%s", *card2)
	if !_card2.Equal(*card2) {
		t.FailNow()
	}

	parsed, err := ParseCards("AhJd6s2c7d")
	t.Logf("parsed cards: %s %s", parsed, err)
	if err != nil {
		t.Errorf("error: %s", err)
	}
}

func TestGenerate(t *testing.T) {
	deck := NewDeck()
	t.Logf("random deck: %s", deck)
}

func TestDiff(t *testing.T) {
	c1, _ := ParseCards("Ah2d3d7c")
	c2, _ := ParseCards("6s2d3d8c")
	result := c1.Diff(c2)
	t.Logf("result=%s", result)
	if result.String() != "Ah7c" {
		t.FailNow()
	}
}

func TestGroupedCards(t *testing.T) {
	c1, _ := ParseCards("AsAhAdAc")
	c2, _ := ParseCards("KsKhKd")
	c3, _ := ParseCards("QsQh")
	c4, _ := ParseCards("Js")
	groups := GroupedCards{c1, c2, c3, c4}
	result := groups.Count()
	t.Logf("result=%s", result)
}
