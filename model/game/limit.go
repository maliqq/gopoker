package game

type Limit string

const (
	FixedLimit Limit = "fixed-limit"
	PotLimit   Limit = "pot-limit"
	NoLimit    Limit = "no-limit"
)

// FIXME
/*
func (limit Limit) BlindsRange(stake *Stake) (float64, float64) {
	sb, bb := stake.Blinds()

	if limit == FixedLimit {
		return sb / 2, sb
	}

	return sb, bb
}
*/

func (limit Limit) RaiseRange(stake *Stake, stackSize float64, potSize float64, bigBets bool) (float64, float64) {
	_, bb := stake.Blinds()

	switch limit {
	case NoLimit:
		return bb, stackSize

	case PotLimit:
		return bb, potSize

	case FixedLimit:
		if bigBets {
			return bb * 2, bb * 2
		}
		return bb, bb
	}

	return 0., 0.
}
