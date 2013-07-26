package protocol

import (
	_ "gopoker/model"
	"gopoker/poker"
)

// hand info
type ShowHand struct {
	Pos        int
	Cards      poker.Cards
	Hand       *poker.Hand
	HandString string
}

// pocket cards show
type ShowCards struct {
	Pos   int
	Cards poker.Cards
	Muck  bool
}

// win info
type Winner struct {
	Pos    int
	Amount float64
}

func NewShowHand(pos int, cards poker.Cards, hand *poker.Hand) *Message {
	return NewMessage(ShowHand{
		Pos:        pos,
		Cards:      cards,
		Hand:       hand,
		HandString: hand.PrintString(),
	})
}

func NewShowCards(pos int, cards poker.Cards) *Message {
	return NewMessage(ShowCards{
		Pos:   pos,
		Cards: cards,
	})
}

func NewMuckCards(pos int, cards *poker.Cards) *Message {
	return NewMessage(ShowCards{
		Pos:   pos,
		Cards: *cards,
		Muck:  true,
	})
}

func NewWinner(pos int, amount float64) *Message {
	return NewMessage(Winner{
		Pos:    pos,
		Amount: amount,
	})
}
