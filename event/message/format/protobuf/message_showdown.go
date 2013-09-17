package protobuf

import (
	"code.google.com/p/goprotobuf/proto"
)

// Cards - byte array
type Cards []byte

// NotifyShowHand - notify show hand
func NewShowHand(pos int, player *string, cards Cards, hand *Hand, handStr string) *Message {
	return &Message{
		ShowHand: &ShowHand{
			Pos:        proto.Int32(int32(pos)),
			Player:     player,
			Cards:      cards,
			Hand:       hand,
			HandString: proto.String(handStr),
		},
	}
}

// NotifyShowCards - notify show cards
func NewShowCards(pos int, player *string, cards Cards) *Message {
	return &Message{
		ShowCards: &ShowCards{
			Pos:    proto.Int32(int32(pos)),
			Player: player,
			Cards:  cards,
		},
	}
}

// NotifyMuckCards - notify muck cards
func NewMuckCards(pos int, player *string, cards Cards) *Message {
	return &Message{
		ShowCards: &ShowCards{
			Pos:    proto.Int32(int32(pos)),
			Player: player,
			Cards:  cards,
			Muck:   proto.Bool(true),
		},
	}
}

// NotifyWinner - notify new winner
func NewWinner(pos int, player *string, amount float64) *Message {
	return &Message{
		Winner: &Winner{
			Pos:    proto.Int32(int32(pos)),
			Player: player,
			Amount: proto.Float64(amount),
		},
	}
}
