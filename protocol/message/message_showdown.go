package message

import (
	"code.google.com/p/goprotobuf/proto"
)

type Cards []byte

func NewShowHand(pos int, cards Cards, hand *Hand, handStr string) *Message {
	return NewMessage(ShowHand{
		Pos:        proto.Int32(int32(pos)),
		Cards:      cards,
		Hand:       hand,
		HandString: proto.String(handStr),
	})
}

func NewShowCards(pos int, cards Cards) *Message {
	return NewMessage(ShowCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards,
	})
}

func NewMuckCards(pos int, cards Cards) *Message {
	return NewMessage(ShowCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards,
		Muck:  proto.Bool(true),
	})
}

func NewWinner(pos int, amount float64) *Message {
	return NewMessage(Winner{
		Pos:    proto.Int32(int32(pos)),
		Amount: proto.Float64(amount),
	})
}
