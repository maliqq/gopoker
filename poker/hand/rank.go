package hand

type Rank string

const (
	StraightFlush Rank = "StraightFlush"
	FourKind      Rank = "FourKind"
	FullHouse     Rank = "FullHouse"
	Flush         Rank = "Flush"
	Straight      Rank = "Straight"
	ThreeKind     Rank = "ThreeKind"
	TwoPair       Rank = "TwoPair"
	OnePair       Rank = "OnePair"
	HighCard      Rank = "HighCard"
	BadugiFour    Rank = "BadugiFour"
	BadugiThree   Rank = "BadugiThree"
	BadugiTwo     Rank = "BadugiTwo"
	BadugiOne     Rank = "BadugiOne"
	CompleteLow   Rank = "CompleteLow"
	IncompleteLow Rank = "IncompleteLow"
)

var (
	Ranks = map[Rank]int{
		StraightFlush: 0,
		FourKind:      1,
		FullHouse:     2,
		Flush:         3,
		Straight:      4,
		ThreeKind:     5,
		TwoPair:       6,
		OnePair:       7,
		HighCard:      8,

		BadugiFour:  0,
		BadugiThree: 1,
		BadugiTwo:   2,
		BadugiOne:   3,

		CompleteLow:   0,
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
