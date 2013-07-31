package message

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model/deal"
)

func NewDealPocket(pos int, cards Cards, dealingType deal.Type) *Message {
	return NewMessage(DealCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards,
		Type:  DealType(DealType_value[string(dealingType)]).Enum(),
	})
}

func NewDealShared(cards Cards, dealingType deal.Type) *Message {
	return NewMessage(DealCards{
		Cards: cards,
		Type:  DealType(DealType_value[string(dealingType)]).Enum(),
	})
}

func NewRequireDiscard(req *RequireDiscard) *Message {
	return NewMessage(*req)
}

func NewDiscarded(pos int, cardsNum int) *Message {
	return NewMessage(Discarded{
		Pos: proto.Int32(int32(pos)),
		Num: proto.Int32(int32(cardsNum)),
	})
}

func NewDiscardCards(pos int, cards Cards) *Message {
	return NewMessage(DiscardCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards,
	})
}
