package message

import (
	"code.google.com/p/goprotobuf/proto"
)

// Cards - byte array
type Cards []byte

// NewShowHand - notify show hand
func NewShowHand(pos int, player *string, cards Cards, hand *Hand, handStr string) *Message {
	return NewMessage(ShowHand{
		Pos:        proto.Int32(int32(pos)),
		Player:     player,
		Cards:      cards,
		Hand:       hand,
		HandString: proto.String(handStr),
	})
}

// NewShowCards - notify show cards
func NewShowCards(pos int, player *string, cards Cards) *Message {
	return NewMessage(ShowCards{
		Pos:    proto.Int32(int32(pos)),
		Player: player,
		Cards:  cards,
	})
}

// NewMuckCards - notify muck cards
func NewMuckCards(pos int, player *string, cards Cards) *Message {
	return NewMessage(ShowCards{
		Pos:    proto.Int32(int32(pos)),
		Player: player,
		Cards:  cards,
		Muck:   proto.Bool(true),
	})
}

// NewWinner - notify new winner
func NewWinner(pos int, player *string, amount float64) *Message {
	return NewMessage(Winner{
		Pos:    proto.Int32(int32(pos)),
		Player: player,
		Amount: proto.Float64(amount),
	})
}
