package poker

import (
	"errors"
	"fmt"
	"gopoker/poker/hand"
)

type Ranking string

const (
	High       Ranking = "high"
	Badugi     Ranking = "badugi"
	AceFive    Ranking = "ace-five"
	AceFive8   Ranking = "ace-five8"
	AceSix     Ranking = "ace-six"
	DeuceSeven Ranking = "deuce-seven"
)

type RankFunc func(*PocketCards) *Hand

type Ranker struct {
	rank     hand.Rank
	rankFunc RankFunc
}

func (r Ranking) Detect(cards *Cards) (*Hand, error) {
	switch r {
	case High:
		return isHigh(cards)

	case Badugi:
		return isBadugi(cards)

	case AceFive:
		return isAceFive(cards)

	case AceFive8:
		return isAceFive8(cards)

	case AceSix:
		return isAceSix(cards)

	case DeuceSeven:
		return isDeuceSeven(cards)
	}
	return nil, errors.New(fmt.Sprintf("unknown ranking: %s", r))
}
