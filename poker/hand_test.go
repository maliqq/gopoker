package poker

import (
	"testing"
)

import (
	"gopoker/poker/hand"
)

func TestHandHumanString(t *testing.T) {
	high, _ := ParseCards("8dTh")
	value, _ := ParseCards("2d5s7hJdQd")
	h := Hand{
		High:  *high,
		Value: *value,
	}

	for rank := hand.StraightFlush; rank <= hand.HighCard; rank++ {
		h.Rank = rank
		t.Logf(h.HumanString())
	}
}
