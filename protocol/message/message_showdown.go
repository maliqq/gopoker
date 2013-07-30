package message

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/poker"
)

func NewShowHand(pos int, cards poker.Cards, hand *poker.Hand) *Message {
	return NewMessage(ShowHand{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards.Binary(),
		Hand: &Hand{
			Rank:   Rank(Rank_value[string(hand.Rank)]).Enum(),
			High:   hand.High.Binary(),
			Value:  hand.Value.Binary(),
			Kicker: hand.Kicker.Binary(),
		},
		HandString: proto.String(hand.PrintString()),
	})
}

func NewShowCards(pos int, cards poker.Cards) *Message {
	return NewMessage(ShowCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards.Binary(),
	})
}

func NewMuckCards(pos int, cards poker.Cards) *Message {
	return NewMessage(ShowCards{
		Pos:   proto.Int32(int32(pos)),
		Cards: cards.Binary(),
		Muck:  proto.Bool(true),
	})
}

func NewWinner(pos int, amount float64) *Message {
	return NewMessage(Winner{
		Pos:    proto.Int32(int32(pos)),
		Amount: proto.Float64(amount),
	})
}
