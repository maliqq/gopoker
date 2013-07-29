package protocol

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/poker"
)

func NewRequireDiscard(req *RequireDiscard) *Message {
	return NewMessage(*req)
}

func NewDiscarded(pos int, cardsNum int) *Message {
	return NewMessage(Discarded{
		Pos: proto.Int32(int32(pos)),
		Num: proto.Int32(int32(cardsNum)),
	})
}

func NewDiscardCards(pos int, cards poker.Cards) *Message {
	return NewMessage(DiscardCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards.Binary(),
	})
}
