package bet

type Type string

const (
	Ante       Type = "Ante"
	BringIn    Type = "BringIn"
	SmallBlind Type = "SmallBlind"
	BigBlind   Type = "BigBlind"
	GuestBlind Type = "GuestBlind"
	Straddle   Type = "Straddle"

	Raise Type = "Raise"
	Call  Type = "Call"

	Check Type = "Check"
	Fold  Type = "Fold"

	Discard  Type = "Discard"
	StandPat Type = "StandPat"

	Show Type = "Show"
	Muck Type = "Muck"
)
