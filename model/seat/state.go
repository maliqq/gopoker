package seat

// State - seat state
type State string

// Seat states
const (
	Empty State = "Empty" // Empty - empty seat
	Taken State = "Taken" // Taken - reserved seat
	Ready State = "Ready" // Ready - ready to play seat

	WaitBB State = "WaitBB" // WaitBB - waiting big blind
	PostBB State = "PostBB" // PostBB - posting big blind after waiting

	Play  State = "Play"  // Play - active seat
	Bet   State = "Bet"   // Bet - have bet in current betting
	AllIn State = "AllIn" // AllIn - gone all-in in current betting
	Fold  State = "Fold"  // Fold - folded in current betting

	Auto State = "Auto" // Auto - autoplay (check/fold strategy)
	Kick State = "Kick" // Kick - kicked out seat

	Away State = "Away" // Away - gone away
	Idle State = "Idle" // Idle - idle seat
)
