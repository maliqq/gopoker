package hand

type Rank string

const (
	StraightFlush Rank = "straight-flush"
	FourKind Rank =      "four-kind"
	FullHouse Rank =     "full-house"
	Flush Rank =         "flush"
	Straight Rank =      "straight"
	ThreeKind Rank =     "three-kind"
	TwoPair Rank =       "two-pair"
	OnePair Rank =       "one-pair"
	HighCard Rank =      "high-card"
	BadugiFour Rank =    "badugi-four"
	BadugiThree Rank =   "badugi-three"
	BadugiTwo Rank =     "badugi-two"
	BadugiOne Rank =     "badugi-one"
	CompleteLow Rank =   "complete-low"
	IncompleteLow Rank = "incomplete-low"
)

var (
	Ranks = map[Rank]int {
		StraightFlush: 0,
		FourKind: 1,
		FullHouse: 2,
		Flush: 3,
		Straight: 4,
		ThreeKind: 5,
		TwoPair: 6,
		OnePair: 7,
		HighCard: 8,

		BadugiFour: 0,
		BadugiThree: 1,
		BadugiTwo: 2,
		BadugiOne: 3,

		CompleteLow: 0,
		IncompleteLow: 1,
	}
)

func (r1 Rank) Compare(r2 Rank) int {
	a := Ranks[r1]
	b := Ranks[r2]

	if a > b {
		return -1
	}
	if a == b {
		return 0
	}

	return 1
}
