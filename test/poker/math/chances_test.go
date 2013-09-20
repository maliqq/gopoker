package math

import (
	"testing"
)

import (
	"gopoker/poker"
)

func TestAgainstN(t *testing.T) {
	against5 := ChancesAgainstN{OpponentsNum: 5, SamplesNum: 1000}

	chances := against5.WithBoard(poker.StringCards("AdAh"), poker.StringCards("As4d5d"))

	if chances.Total() != 5000 {
		t.Fatalf("chances=%s", chances)
	}
}
