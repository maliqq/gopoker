package poker

import (
	"testing"
)

import (
	"gopoker/poker/hand"
	"gopoker/poker/ranking"
)

func TestLow(t *testing.T) {
	cards, _ := ParseCards("Ad2c3c4s5s6h")
	h, _ := Detect[ranking.AceFive](cards)

	t.Logf("AceFive=%s", h)
	if h.Rank != hand.CompleteLow {
		t.FailNow()
	}

	cards, _ = ParseCards("2c3c4sAd8d")
	h, _ = Detect[ranking.AceFive8](cards)

	t.Logf("AceFive8=%s", h)
	if h.Rank != hand.CompleteLow {
		t.FailNow()
	}

	cards, _ = ParseCards("Ad2c3c4s6h8d")
	h, _ = Detect[ranking.AceSix](cards)

	t.Logf("AceSix=%s", h)
	if h.Rank != hand.CompleteLow {
		t.FailNow()
	}

	cards, _ = ParseCards("2c3c4s5sAd7c")
	h, _ = Detect[ranking.DeuceSeven](cards)

	t.Logf("DeuceSeven=%s", h)
	if h.Rank != hand.CompleteLow {
		t.FailNow()
	}
}
