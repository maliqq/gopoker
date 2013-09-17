package protobuf

import (
	"code.google.com/p/goprotobuf/proto"
)

// NotifyMoveButton - notify move button
func NewMoveButton(pos int) *Message {
	return &Message{
		MoveButton: &MoveButton{
			Pos: proto.Int32(int32(pos)),
		},
	}
}

// NotifyJoinTable - notify join table
func NewJoinTable(player *string, pos int, amount float64) *Message {
	return &Message{
		JoinTable: &JoinTable{
			Player: player,
			Pos:    proto.Int32(int32(pos)),
			Amount: proto.Float64(amount),
		},
	}
}

// NotifyLeaveTable - notify leave table
func NewLeaveTable(player *string) *Message {
	return &Message{
		LeaveTable: &LeaveTable{
			Player: player,
		},
	}
}

func NewSitOut(pos int) *Message {
	return &Message{
		SitOut: &SitOut{
			Pos: proto.Int32(int32(pos)),
		},
	}
}

func NewComeBack(pos int) *Message {
	return &Message{
		ComeBack: &ComeBack{
			Pos: proto.Int32(int32(pos)),
		},
	}
}
