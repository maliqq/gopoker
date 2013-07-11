package poker

import (
	"testing"
)

import (
	"gopoker/poker/ranking"
)

func TestLow(t *testing.T) {
	cards, _ := ParseCards("Ad2c3c4s5s6h")
	hand, _ := Detect[ranking.AceFive](cards)

	t.Logf("AceFive=%s", hand)
	t.FailNow()
}
