package state

type Type string

const (
	Waiting Type = "waiting"
	Active  Type = "active"
	Paused  Type = "paused"
	Closed  Type = "closed"
)
