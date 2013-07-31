package message

import (
	"code.google.com/p/goprotobuf/proto"
)

func NewMoveButton(pos int) *Message {
	return NewMessage(MoveButton{
		Pos: proto.Int32(int32(pos)),
	})
}

func NewJoinTable(player string, pos int, amount float64) *Message {
	return NewMessage(JoinTable{
		Player: proto.String(string(player)),
		Pos:    proto.Int32(int32(pos)),
		Amount: proto.Float64(amount),
	})
}

func NewLeaveTable(player string) *Message {
	return NewMessage(LeaveTable{
		Player: proto.String(string(player)),
	})
}
