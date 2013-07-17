package util

import (
	"testing"
)

import (
	"gopoker/poker"
)

func TestEquity(t *testing.T) {
	for i := 0; i <= 10; i++ {
		c := poker.GenerateCards(2)
		t.Logf("equity of %s = %.2f", c, Equity(c))
	}
	t.FailNow()
}
