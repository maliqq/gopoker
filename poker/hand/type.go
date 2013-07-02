package hand

import (
	"strconv"
)

type Rank int

const (
	StraightFlush Rank = iota
	FourKind
	FullHouse
	Flush
	Straight
	ThreeKind
	TwoPair
	OnePair
	HighCard

	BadugiFour
	BadugiThree
	BadugiTwo
	BadugiOne

	CompleteLow
	IncompleteLow
)

var names = map[Rank]string{
	StraightFlush: "straight-flush",
	FourKind:      "four-kind",
	FullHouse:     "full-house",
	Flush:         "flush",
	Straight:      "straight",
	ThreeKind:     "three-kind",
	TwoPair:       "two-pair",
	OnePair:       "one-pair",
	HighCard:      "high-card",
	BadugiFour:    "badugi-four",
	BadugiThree:   "badugi-three",
	BadugiTwo:     "badugi-two",
	BadugiOne:     "badugi-one",
	CompleteLow:   "complete-low",
	IncompleteLow: "incomplete-low",
}

func (r Rank) String() string {
	return names[r]
}

func (r Rank) MarshalJSON() ([]byte, error) {
	s := r.String()
	return []byte(strconv.Quote(s)), nil
}

func (a Rank) Compare(b Rank) int {
	if a > b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}
