package poker

import (
	"testing"
)

func TestHighFast(t *testing.T) {
	cards, _ := ParseCards("AdKdQdJdTd7s8s")
	c := cards.Uint64()

	t.Logf("cards.Uint64=%d", c)

	InitFast()

	rank := doRank(c)
	t.Logf("rank=%d", int(rank) >> 12)

	t.FailNow()
}
