package bet

// Type - bet type
type Type string

const (
	// Ante - ante bet
	Ante Type = "Ante"
	// BringIn - bring in forced bet
	BringIn Type = "BringIn"
	// SmallBlind - small blind forced bet
	SmallBlind Type = "SmallBlind"
	// BigBlind - big blind forced bet
	BigBlind Type = "BigBlind"
	// GuestBlind - guest blind forced bet
	GuestBlind Type = "GuestBlind"
	// Straddle - straddle forced bet
	Straddle Type = "Straddle"

	// Raise - raise active bet
	Raise Type = "Raise"
	// Call - call active bet
	Call Type = "Call"

	// Check - check passive bet
	Check Type = "Check"
	// Fold - fold passive bet
	Fold Type = "Fold"

	// Discard - discard cards
	Discard Type = "Discard"
	// StandPat - stand pat
	StandPat Type = "StandPat"

	// Show - show cards
	Show Type = "Show"
	// Muck - muck cards
	Muck Type = "Muck"
)
