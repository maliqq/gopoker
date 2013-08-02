package hand

// Rank - cards poker rank
type Rank string

const (
	// StraightFlush - straight flush rank
	StraightFlush Rank = "StraightFlush"
	// FourKind - four of a kind rank
	FourKind Rank = "FourKind"
	// FullHouse - full house rank
	FullHouse Rank = "FullHouse"
	// Flush - flush rank
	Flush Rank = "Flush"
	// Straight - straight rank
	Straight Rank = "Straight"
	// ThreeKind - three of a kind rank
	ThreeKind Rank = "ThreeKind"
	// TwoPair - two pair rank
	TwoPair Rank = "TwoPair"
	// OnePair - one pair rank
	OnePair Rank = "OnePair"
	// HighCard - high card rank
	HighCard Rank = "HighCard"
	// BadugiFour - badugi four cards rank
	BadugiFour Rank = "BadugiFour"
	// BadugiThree - badugi three cards rank
	BadugiThree Rank = "BadugiThree"
	// BadugiTwo - badugi two cards rank
	BadugiTwo Rank = "BadugiTwo"
	// BadugiOne - badugi one card rank
	BadugiOne Rank = "BadugiOne"
	// CompleteLow - complete low rank
	CompleteLow Rank = "CompleteLow"
	// IncompleteLow - incomplete low rank
	IncompleteLow Rank = "IncompleteLow"
)

var ranks = map[Rank]int{
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

// Compare - compare two ranks
func (r1 Rank) Compare(r2 Rank) int {
	a := ranks[r1]
	b := ranks[r2]

	if a > b {
		return -1
	}
	if a == b {
		return 0
	}

	return 1
}
