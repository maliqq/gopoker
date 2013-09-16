package model

import (
	"testing"
)

func TestDealer(t *testing.T) {
	dealer := NewDealer()

	dealer.burn(2)
	cards := dealer.Deal(2)
	t.Logf(cards.String())
	discarded := dealer.Discard(cards)
	t.Logf(discarded.String())

	if cards.String() == discarded.String() {
		t.FailNow()
	}

	shared := dealer.Share(2)
	t.Logf(shared.String())
}
