package calc

import (
	"testing"
)

import (
	"gopoker/poker"
)

func TestSklanskyMalmuthGroup(t *testing.T) {
	cards := poker.StringCards("AhAd")
	if g := SklanskyMalmuthGroup(cards[0], cards[1]); g != 1 {
		t.FailNow()
	}

	cards = poker.StringCards("As2h")
	if g := SklanskyMalmuthGroup(cards[0], cards[1]); g != 9 {
		t.FailNow()
	}

	cards = poker.StringCards("As2s")
	if g := SklanskyMalmuthGroup(cards[0], cards[1]); g != 5 {
		t.FailNow()
	}
}
