package seat

// State - seat state
type State string

const (
	// Empty - empty seat
	Empty State = "Empty"
	// Taken - reserved seat
	Taken State = "Taken"
	// Ready - ready to play seat
	Ready State = "Ready"

	// WaitBB - waiting big blind
	WaitBB State = "WaitBB"
	// PostBB - posting big blind after waiting
	PostBB State = "PostBB"

	// Play - active seat
	Play State = "Play"
	// Bet - have bet in current betting
	Bet State = "Bet"
	// AllIn - gone all-in in current betting
	AllIn State = "AllIn"
	// Fold - folded in current betting
	Fold State = "Fold"

	// Auto - autoplay (check/fold strategy)
	Auto State = "Auto"
	// Kick - kicked out seat
	Kick State = "Kick"

	// Away - gone away
	Away State = "Away"
	// Idle - idle seat
	Idle State = "Idle"
)
