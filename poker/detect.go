package poker

import (
	"gopoker/poker/hand"
)

type detectFunc func(*Cards) (*Hand, error)

var Detect = map[hand.Ranking]detectFunc{
	hand.High:       isHigh,
	hand.Badugi:     isBadugi,
	hand.AceFive:    isAceFive,
	hand.AceFive8:   isAceFive8,
	hand.AceSix:     isAceSix,
	hand.DeuceSeven: isDeuceSeven,
}
