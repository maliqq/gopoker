package game

type Limit string

const (
	FixedLimit Limit = "fixed-limit"
	PotLimit   Limit = "pot-limit"
	NoLimit    Limit = "no-limit"
)

func (limit Limit) BlindsRange(stake *Stake) (float64, float64) {
	sb, bb := stake.Blinds()

	if limit == FixedLimit {
		return sb / 2, sb
	}

	return sb, bb
}

func (limit Limit) RaiseRange(stake *Stake, stackSize float64, potSize float64, bigBets bool) (float64, float64) {
	sb, bb := stake.Blinds()

	switch limit {
	case NoLimit:
		return bb, stackSize

	case PotLimit:
		return bb, potSize

	case FixedLimit:
		if bigBets {
			return bb, bb * 2
		}
		return sb, bb
	}

	return 0., 0.
}
