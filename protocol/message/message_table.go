package message

import (
	"code.google.com/p/goprotobuf/proto"
)

// NotifyMoveButton - notify move button
func NotifyMoveButton(pos int) *Message {
	return NewMessage(MoveButton{
		Pos: proto.Int32(int32(pos)),
	})
}

// NotifyJoinTable - notify join table
func NotifyJoinTable(player string, pos int, amount float64) *Message {
	return NewMessage(JoinTable{
		Player: proto.String(string(player)),
		Pos:    proto.Int32(int32(pos)),
		Amount: proto.Float64(amount),
	})
}

// NotifyLeaveTable - notify leave table
func NotifyLeaveTable(player string) *Message {
	return NewMessage(LeaveTable{
		Player: proto.String(string(player)),
	})
}
