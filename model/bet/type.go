package bet

type Type string

const (
	Ante       Type = "ante"
	BringIn    Type = "bring-in"
	SmallBlind Type = "small-blind"
	BigBlind   Type = "big-blind"
	GuestBlind Type = "guest-blind"
	Straddle   Type = "straddle"

	Raise Type = "raise"
	Call  Type = "call"

	Check Type = "check"
	Fold  Type = "fold"

	Discard  Type = "discard"
	StandPat Type = "stand-pat"

	Show Type = "show"
	Muck Type = "muck"
)
