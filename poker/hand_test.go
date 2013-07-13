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

	h.Rank = hand.StraightFlush
	t.Logf(h.HumanString())

	h.Rank = hand.FourKind
	t.Logf(h.HumanString())

	h.Rank = hand.FullHouse
	t.Logf(h.HumanString())

	h.Rank = hand.Flush
	t.Logf(h.HumanString())

	h.Rank = hand.Straight
	t.Logf(h.HumanString())

	h.Rank = hand.ThreeKind
	t.Logf(h.HumanString())

	h.Rank = hand.TwoPair
	t.Logf(h.HumanString())

	h.Rank = hand.OnePair
	t.Logf(h.HumanString())

	h.Rank = hand.HighCard
	t.Logf(h.HumanString())

	t.FailNow()
}
