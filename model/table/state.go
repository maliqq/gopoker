package table

type State int

const (
	Waiting State = iota
	Active
	Paused
	Closed
)
