package poker

import (
	"sort"
	"testing"
)

func TestSortByKind(t *testing.T) {
	parsed, _ := ParseCards("7s3s6s4s")
	sort.Sort(ByKind{parsed, AceHigh})

	t.Logf("sorted by kind=%s", parsed)
}

func TestSortBySuit(t *testing.T) {
	parsed, _ := ParseCards("3d3s4d4c4s")
	sort.Sort(BySuit{parsed})

	t.Logf("sorted by suit=%s", parsed)
}
