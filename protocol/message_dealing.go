package protocol

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model/deal"
	"gopoker/poker"
)

func NewDealPocket(pos int, cards poker.Cards, dealingType deal.Type) *Message {
	return NewMessage(DealCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards.Binary(),
		Type:  DealType(DealType_value[string(dealingType)]).Enum(),
	})
}

func NewDealShared(cards poker.Cards, dealingType deal.Type) *Message {
	return NewMessage(DealCards{
		Cards: cards.Binary(),
		Type:  DealType(DealType_value[string(dealingType)]).Enum(),
	})
}
