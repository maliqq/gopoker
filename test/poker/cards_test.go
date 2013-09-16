package poker

import (
	"encoding/json"
	"testing"
)

func TestCards(t *testing.T) {
	deck := NewDeck()

	cards := deck[0:3]
	t.Logf("cards=%s", cards)
	t.Logf("combine(2)=%s", cards.Combine(2))
	t.Logf("combine(3)=%s", cards.Combine(3))

	cards = deck[0:5]
	t.Logf("cards=%s", cards)
	t.Logf("combine(2)=%s", cards.Combine(2))
	t.Logf("combine(3)=%s", cards.Combine(3))

	//t.FailNow()
}

func TestBinary(t *testing.T) {
	c, _ := ParseCards("AhKdJs8h")

	b, _ := json.Marshal(c)
	t.Logf("raw=%#v json=%s", c.Binary(), b)

	c = Cards{nil, &Card{0, 0}, &Card{2, 3}}

	b, _ = json.Marshal(c)

	t.Logf("raw=%#v json=%s", c.Binary(), b)
}
