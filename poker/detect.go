package poker

import (
	"gopoker/poker/ranking"
)

type detectFunc func(*Cards) (*Hand, error)

var (
	Detect = map[ranking.Type]detectFunc{
		ranking.High: func(cards *Cards) (*Hand, error) {
			return isHigh(cards)
		},

		ranking.Badugi: func(cards *Cards) (*Hand, error) {
			return isBadugi(cards)
		},

		ranking.AceFive: func(cards *Cards) (*Hand, error) {
			return isAceFive(cards)
		},

		ranking.AceFive8: func(cards *Cards) (*Hand, error) {
			return isAceFive8(cards)
		},

		ranking.AceSix: func(cards *Cards) (*Hand, error) {
			return isAceSix(cards)
		},

		ranking.DeuceSeven: func(cards *Cards) (*Hand, error) {
			return isDeuceSeven(cards)
		},
	}
)
