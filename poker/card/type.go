package card

type Tuple struct {
	Kind
	Suit
}

const (
	CardsNum = KindsNum * SuitsNum
)

var (
	Masks = AllMasks()
)

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

func AllMasks() []uint64 {
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
