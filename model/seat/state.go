package seat

type State string

const (
	Empty State = "empty"
	Taken State = "taken"
	Ready State = "ready"

	WaitBB State = "wait-bb"
	PostBB State = "post-bb"

	Play  State = "play"
	Bet   State = "bet"
	AllIn State = "all-in"
	Fold  State = "fold"

	Auto State = "auto"
	Kick State = "kick"

	Away State = "away"
	Idle State = "idle"
)
