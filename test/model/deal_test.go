package model

import (
	"testing"
)

func TestDeal(t *testing.T) {
	d := NewDeal()
	a := Player("A")
	b := Player("B")
	d.DealPocket(a, 2)

	cards1 := d.Pocket(a)
	t.Logf("cards1 = %s", cards1)
	if len(cards1) != 2 {
		t.FailNow()
	}

	cards2 := d.DealPocket(b, 4)
	t.Logf("cards2 = %s", cards2)
	if len(cards2) != 4 {
		t.FailNow()
	}

	cards3 := d.Discard(b, cards2)
	t.Logf("cards3 = %s", cards3)
	if len(cards3) != 4 {
		t.FailNow()
	}

	cards4 := d.DealBoard(5)
	t.Logf("cards4 = %s", cards4)
	if len(cards4) != 5 {
		t.FailNow()
	}
}
