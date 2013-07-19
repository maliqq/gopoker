package util

import (
	"testing"
)

import (
	"gopoker/poker"
)

func TestEquity(t *testing.T) {
	for kind := 0; kind < 13; kind++ {
		c1, _ := poker.MakeCard(kind, 0)
		c2, _ := poker.MakeCard(kind, 1)
		c := poker.Cards{*c1, *c2}
		t.Logf("equity of %s = %.4f", c, Equity(c))
	}
	t.FailNow()
}
