package protocol

import (
	"gopoker/poker"
)

type RequireDiscard struct {
	Pos int
}

type Discarded struct {
	Pos int
	Num int
}

type DiscardCards struct {
	Pos   int
	Cards poker.Cards
}

func NewRequireDiscard(req *RequireDiscard) *Message {
	return NewMessage(*req)
}

func NewDiscarded(pos int, cardsNum int) *Message {
	return NewMessage(
		Discarded{
			Pos: pos,
			Num: cardsNum,
		},
	)
}

func NewDiscardCards(pos int, cards poker.Cards) *Message {
	return NewMessage(
		DiscardCards{
			Pos:   pos,
			Cards: cards,
		},
	)
}
