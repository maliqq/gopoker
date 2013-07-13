package bet

type ForcedBet string
type ActiveBet string
type PassiveBet string
type AnyBet string

type Type interface {
	Value() AnyBet
}

const (
	Ante       ForcedBet = "ante"
	BringIn    ForcedBet = "bring-in"
	SmallBlind ForcedBet = "small-blind"
	BigBlind   ForcedBet = "big-blind"
	GuestBlind ForcedBet = "guest-blind"
	Straddle   ForcedBet = "straddle"

	Raise ActiveBet = "raise"
	Call  ActiveBet = "call"

	Check PassiveBet = "check"
	Fold  PassiveBet = "fold"

	Discard  AnyBet = "discard"
	StandPat AnyBet = "stand-pat"

	Show AnyBet = "show"
	Muck AnyBet = "muck"
)

func (any AnyBet) Value() AnyBet {
	return any
}

func (forced ForcedBet) Value() AnyBet {
	return AnyBet(string(forced))
}

func (active ActiveBet) Value() AnyBet {
	return AnyBet(string(active))
}

func (passive PassiveBet) Value() AnyBet {
	return AnyBet(string(passive))
}
