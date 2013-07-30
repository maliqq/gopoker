package seat

type State string

const (
	Empty State = "Empty"
	Taken State = "Taken"
	Ready State = "Ready"

	WaitBB State = "WaitBB"
	PostBB State = "PostBB"

	Play  State = "Play"
	Bet   State = "Bet"
	AllIn State = "AllIn"
	Fold  State = "Fold"

	Auto State = "Auto"
	Kick State = "Kick"

	Away State = "Away"
	Idle State = "Idle"
)
