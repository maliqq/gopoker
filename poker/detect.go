package poker

import (
	"gopoker/poker/ranking"
)

type detectFunc func(*Cards) (*Hand, error)

var Detect = map[ranking.Type]detectFunc{
	ranking.High:       isHigh,
	ranking.Badugi:     isBadugi,
	ranking.AceFive:    isAceFive,
	ranking.AceFive8:   isAceFive8,
	ranking.AceSix:     isAceSix,
	ranking.DeuceSeven: isDeuceSeven,
}
