package bet

// Type - bet type
type Type string

// Bet types
const (
	Ante       Type = "Ante"       // Ante - ante bet
	BringIn    Type = "BringIn"    // BringIn - bring in forced bet
	SmallBlind Type = "SmallBlind" // SmallBlind - small blind forced bet
	BigBlind   Type = "BigBlind"   // BigBlind - big blind forced bet
	GuestBlind Type = "GuestBlind" // GuestBlind - guest blind forced bet
	Straddle   Type = "Straddle"   // Straddle - straddle forced bet

	Raise Type = "Raise" // Raise - raise active bet
	Call  Type = "Call"  // Call - call active bet
	Check Type = "Check" // Check - check passive bet
	Fold  Type = "Fold"  // Fold - fold passive bet

	Discard  Type = "Discard"  // Discard - discard cards
	StandPat Type = "StandPat" // StandPat - stand pat

	Show Type = "Show" // Show - show cards
	Muck Type = "Muck" // Muck - muck cards
)
