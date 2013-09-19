package message

import (
	"gopoker/model"
)

// MoveButton
type MoveButton struct {
	Pos int
}

func (MoveButton) EventMessage() {}

// JoinTable
type JoinTable struct {
	Player model.Player
	Pos    int
	Amount float64
}

func (JoinTable) EventMessage() {}

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
type LeaveTable struct {
	Player model.Player
}

func (LeaveTable) EventMessage() {}
