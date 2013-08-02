package message

import (
	"code.google.com/p/goprotobuf/proto"
)

// NewMoveButton - notify move button
func NewMoveButton(pos int) *Message {
	return NewMessage(MoveButton{
		Pos: proto.Int32(int32(pos)),
	})
}

// NewJoinTable - notify join table
func NewJoinTable(player string, pos int, amount float64) *Message {
	return NewMessage(JoinTable{
		Player: proto.String(string(player)),
		Pos:    proto.Int32(int32(pos)),
		Amount: proto.Float64(amount),
	})
}

// NewLeaveTable - notify leave table
func NewLeaveTable(player string) *Message {
	return NewMessage(LeaveTable{
		Player: proto.String(string(player)),
	})
}
