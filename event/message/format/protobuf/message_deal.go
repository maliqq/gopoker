package protobuf

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model/deal"
)

// NotifyDealCards - notify new pocket cards
func NewDealCards(pos int, cards Cards, dealingType deal.Type) *Message {
	return &Message{
		DealCards: &DealCards{
			Pos:   proto.Int32(int32(pos)),
			Cards: cards,
			Type:  DealType(DealType_value[string(dealingType)]).Enum(),
		},
	}
}

// NotifyRequireDiscard - notify require discard
func NewRequireDiscard(pos int) *Message {
	return &Message{
		RequireDiscard: &RequireDiscard{
			Pos: proto.Int32(int32(pos)),
		},
	}
}

// NotifyDiscarded - notify discarded action
func NewDiscarded(pos int, cardsNum int) *Message {
	return &Message{
		Discarded: &Discarded{
			Pos: proto.Int32(int32(pos)),
			Num: proto.Int32(int32(cardsNum)),
		},
	}
}

// NotifyDiscardCards - notify exchanged cards
func NewDiscardCards(pos int, cards Cards) *Message {
	return &Message{
		DiscardCards: &DiscardCards{
			Pos:   proto.Int32(int32(pos)),
			Cards: cards,
		},
	}
}
