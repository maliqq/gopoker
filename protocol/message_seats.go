package protocol

import (
	"gopoker/model/seat"
)

type SeatState struct {
	Pos int
}

type ChangeSeatState struct {
	Pos   int
	State seat.State
}

type SitOut struct {
	Pos int
}

type ComeBack struct {
	Pos int
}

// seat info
type Seat struct {
	State string
	Stack float64
	Bet   float64
}

type SeatStack struct {
	Pos int
}

type ChangeSeatStack struct {
	Pos    int
	Amount float64
}

type AdvanceSeatStack struct {
	Pos    int
	Amount float64
}
