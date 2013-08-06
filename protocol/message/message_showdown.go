package message

import (
	"code.google.com/p/goprotobuf/proto"
)

// Cards - byte array
type Cards []byte

// NotifyShowHand - notify show hand
func NotifyShowHand(pos int, player *string, cards Cards, hand *Hand, handStr string) *Message {
	return NewMessage(ShowHand{
		Pos:        proto.Int32(int32(pos)),
		Player:     player,
		Cards:      cards,
		Hand:       hand,
		HandString: proto.String(handStr),
	})
}

// NotifyShowCards - notify show cards
func NotifyShowCards(pos int, player *string, cards Cards) *Message {
	return NewMessage(ShowCards{
		Pos:    proto.Int32(int32(pos)),
		Player: player,
		Cards:  cards,
	})
}

// NotifyMuckCards - notify muck cards
func NotifyMuckCards(pos int, player *string, cards Cards) *Message {
	return NewMessage(ShowCards{
		Pos:    proto.Int32(int32(pos)),
		Player: player,
		Cards:  cards,
		Muck:   proto.Bool(true),
	})
}

// NotifyWinner - notify new winner
func NotifyWinner(pos int, player *string, amount float64) *Message {
	return NewMessage(Winner{
		Pos:    proto.Int32(int32(pos)),
		Player: player,
		Amount: proto.Float64(amount),
	})
}
