package protocol

import (
	"gopoker/model"
)

type MoveButton struct {
	Pos int
}

type JoinTable struct {
	Player model.Player
	Pos    int
	Amount float64
}

type LeaveTable struct {
	Player model.Player
}

func NewMoveButton(pos int) *Message {
	return NewMessage(MoveButton{
		Pos: pos,
	})
}

func NewJoinTable(player model.Player, pos int, amount float64) *Message {
	return NewMessage(JoinTable{
		Player: player,
		Pos:    pos,
		Amount: amount,
	})
}

func NewLeaveTable(player model.Player) *Message {
	return NewMessage(LeaveTable{
		Player: player,
	})
}

type ChangeTableState struct {
	State string
}
