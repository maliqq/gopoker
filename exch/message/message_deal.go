package message

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model/deal"
)

// NotifyDealPocket - notify new pocket cards
func NotifyDealPocket(pos int, cards Cards, dealingType deal.Type) *Message {
	return NewMessage(DealCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards,
		Type:  DealType(DealType_value[string(dealingType)]).Enum(),
	})
}

// NotifyDealShared - notify new shared cards
func NotifyDealShared(cards Cards, dealingType deal.Type) *Message {
	return NewMessage(DealCards{
		Cards: cards,
		Type:  DealType(DealType_value[string(dealingType)]).Enum(),
	})
}

// NotifyRequireDiscard - notify require discard
func NotifyRequireDiscard(req *RequireDiscard) *Message {
	return NewMessage(*req)
}

// NotifyDiscarded - notify discarded action
func NotifyDiscarded(pos int, cardsNum int) *Message {
	return NewMessage(Discarded{
		Pos: proto.Int32(int32(pos)),
		Num: proto.Int32(int32(cardsNum)),
	})
}

// NotifyDiscardCards - notify exchanged cards
func NotifyDiscardCards(pos int, cards Cards) *Message {
	return NewMessage(DiscardCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards,
	})
}
