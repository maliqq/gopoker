package poker

import (
	"testing"
	"reflect"
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

func TestMask(t *testing.T) {
	table := []uint64{
		0x0001000000000000, 0x0002000000000000, 0x0004000000000000, 0x0008000000000000,
		0x0010000000000000, 0x0020000000000000, 0x0040000000000000, 0x0080000000000000,
		0x0100000000000000, 0x0200000000000000, 0x0400000000000000, 0x0800000000000000,
		0x1000000000000000,
		
		0x0000000100000000, 0x0000000200000000, 0x0000000400000000, 0x0000000800000000,
		0x0000001000000000, 0x0000002000000000, 0x0000004000000000, 0x0000008000000000,
		0x0000010000000000, 0x0000020000000000, 0x0000040000000000, 0x0000080000000000,
		0x0000100000000000,
		
		0x0000000000010000, 0x0000000000020000, 0x0000000000040000, 0x0000000000080000,
		0x0000000000100000, 0x0000000000200000, 0x0000000000400000, 0x0000000000800000,
		0x0000000001000000, 0x0000000002000000, 0x0000000004000000, 0x0000000008000000,
		0x0000000010000000,

		0x0000000000000001, 0x0000000000000002, 0x0000000000000004, 0x0000000000000008,
		0x0000000000000010, 0x0000000000000020, 0x0000000000000040, 0x0000000000000080,
		0x0000000000000100, 0x0000000000000200, 0x0000000000000400, 0x0000000000000800,
		0x0000000000001000,
	}

	if !reflect.DeepEqual(card.Masks, table) {
		t.Fatalf("masks mismatch: \ntable=%#v\nmasks=%#v", table, card.Masks)
	}
}

func TestParse(t *testing.T) {
	card1, _ := ParseCard("2s")
	_card1, _ := NewCard(1)
	t.Logf("card1=%s", card1)
	if !_card1.Equal(card1) {
		t.FailNow()
	}

	card2, _ := ParseCard("Ac")
	_card2, _ := NewCard(52)
	t.Logf("card2=%s", card2)
	if !_card2.Equal(card2) {
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
