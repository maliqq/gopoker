package message

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model/deal"
)

// NewDealPocket - notify new pocket cards
func NewDealPocket(pos int, cards Cards, dealingType deal.Type) *Message {
	return NewMessage(DealCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards,
		Type:  DealType(DealType_value[string(dealingType)]).Enum(),
	})
}

// NewDealShared - notify new shared cards
func NewDealShared(cards Cards, dealingType deal.Type) *Message {
	return NewMessage(DealCards{
		Cards: cards,
		Type:  DealType(DealType_value[string(dealingType)]).Enum(),
	})
}

// NewRequireDiscard - notify require discard
func NewRequireDiscard(req *RequireDiscard) *Message {
	return NewMessage(*req)
}

// NewDiscarded - notify discarded action
func NewDiscarded(pos int, cardsNum int) *Message {
	return NewMessage(Discarded{
		Pos: proto.Int32(int32(pos)),
		Num: proto.Int32(int32(cardsNum)),
	})
}

// NewDiscardCards - notify exchanged cards
func NewDiscardCards(pos int, cards Cards) *Message {
	return NewMessage(DiscardCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards,
	})
}
