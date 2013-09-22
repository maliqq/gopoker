package message

import (
	"gopoker/model"
  "gopoker/model/seat"
)

type Seat struct {
  Pos int
}

func (Seat) EventMessage() {}

type SeatState struct {
  Pos int
  State seat.State
}

func (SeatState) EventMessage() {}

// MoveButton
type MoveButton struct {
	Pos int
}

func (MoveButton) EventMessage() {}

type Join struct {
	Player model.Player
	Pos    int
	Amount float64
}

func (Join) EventMessage() {}

// SitOut
type SitOut struct {
	Pos int
}

func (SitOut) EventMessage() {}

// ComeBack
type ComeBack struct {
	Pos int
}

func (ComeBack) EventMessage() {}

// LeaveTable
type Leave struct {
	Player model.Player
}

func (Leave) EventMessage() {}
