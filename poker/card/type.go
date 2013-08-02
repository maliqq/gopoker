package card

// Tuple - kind and suit pair
type Tuple struct {
	Kind
	Suit
}

const (
	// CardsNum - 52
	CardsNum = KindsNum * SuitsNum
)

var (
	// Masks - bits for all cards
	Masks = allMasks()
)

// AllTuples - all tuples
func AllTuples() []Tuple {
	cards := make([]Tuple, CardsNum)

	k := 0
	for _, kind := range AllKinds() {
		for _, suit := range AllSuits() {
			cards[k] = Tuple{kind, suit}
			k++
		}
	}

	return cards
}

func allMasks() []uint64 {
	masks := make([]uint64, CardsNum)
	i := 0
	for suit := SuitsNum - 1; suit >= 0; suit-- {
		for kind := 0; kind < KindsNum; kind++ {
			masks[i] = uint64((1 << uint(kind) << uint(1<<4*suit)))
			i++
		}
	}
	return masks
}
