package poker

import (
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
