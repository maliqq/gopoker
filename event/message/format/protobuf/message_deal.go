package protobuf

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model/deal"
)

// NotifyDealPocket - notify new pocket cards
func NewDealPocket(pos int, cards Cards, dealingType deal.Type) *Message {
	return &Message{
		Payload: &Payload{
			DealCards: &DealCards{
				Pos:   proto.Int32(int32(pos)),
				Cards: cards,
				Type:  DealType(DealType_value[string(dealingType)]).Enum(),
			},
		},
	}
}

// NotifyDealShared - notify new shared cards
func NewDealShared(cards Cards, dealingType deal.Type) *Message {
	return &Message{
		Payload: &Payload{
			DealCards: &DealCards{
				Cards: cards,
				Type:  DealType(DealType_value[string(dealingType)]).Enum(),
			},
		},
	}
}

// NotifyRequireDiscard - notify require discard
func NewRequireDiscard(req *RequireDiscard) *Message {
	return &Message{
		Payload: &Payload{
			RequireDiscard: req,
		},
	}
}

// NotifyDiscarded - notify discarded action
func NewDiscarded(pos int, cardsNum int) *Message {
	return &Message{
		Payload: &Payload{
			Discarded: &Discarded{
				Pos: proto.Int32(int32(pos)),
				Num: proto.Int32(int32(cardsNum)),
			},
		},
	}
}

// NotifyDiscardCards - notify exchanged cards
func NewDiscardCards(pos int, cards Cards) *Message {
	return &Message{
		Payload: &Payload{
			DiscardCards: &DiscardCards{
				Pos:   proto.Int32(int32(pos)),
				Cards: cards,
			},
		},
	}
}
