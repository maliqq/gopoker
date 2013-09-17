package protobuf

import (
	"code.google.com/p/goprotobuf/proto"
)

// NotifyMoveButton - notify move button
func NewMoveButton(pos int) *Message {
	return &Message{
		Payload: &Payload{
			MoveButton: &MoveButton{
				Pos: proto.Int32(int32(pos)),
			},
		},
	}
}

// NotifyJoinTable - notify join table
func NewJoinTable(player string, pos int, amount float64) *Message {
	return &Message{
		Payload: &Payload{
			JoinTable: &JoinTable{
				Player: proto.String(string(player)),
				Pos:    proto.Int32(int32(pos)),
				Amount: proto.Float64(amount),
			},
		},
	}
}

// NotifyLeaveTable - notify leave table
func NewLeaveTable(player string) *Message {
	return &Message{
		Payload: &Payload{
			LeaveTable: &LeaveTable{
				Player: proto.String(string(player)),
			},
		},
	}
}
