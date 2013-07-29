package protocol

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model"
)

func NewMoveButton(pos int) *Message {
	return NewMessage(MoveButton{
		Pos: proto.Int32(int32(pos)),
	})
}

func NewJoinTable(player model.Player, pos int, amount float64) *Message {
	return NewMessage(JoinTable{
		Player: proto.String(string(player)),
		Pos:    proto.Int32(int32(pos)),
		Amount: proto.Float64(amount),
	})
}

func NewLeaveTable(player model.Player) *Message {
	return NewMessage(LeaveTable{
		Player: proto.String(string(player)),
	})
}
