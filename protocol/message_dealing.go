package protocol

import (
	"gopoker/model/deal"
	"gopoker/poker"
)

type DealCards struct {
	Pos   int
	Cards poker.Cards
	Type  deal.Type
}

func NewDealPocket(pos int, cards poker.Cards, dealingType deal.Type) *Message {
	return NewMessage(
		DealCards{
			Pos:   pos,
			Cards: cards,
			Type:  dealingType,
		},
	)
}

func NewDealShared(cards poker.Cards, dealingType deal.Type) *Message {
	return NewMessage(
		DealCards{
			Cards: cards,
			Type:  dealingType,
		},
	)
}
