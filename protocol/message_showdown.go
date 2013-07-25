package protocol

import (
	"gopoker/model"
	"gopoker/poker"
)

// hand info
type ShowHand struct {
	Pos   int
	Cards poker.Cards
	Hand  *poker.Hand
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
	Player model.Player
	Amount float64
}

func NewShowHand(pos int, cards poker.Cards, hand *poker.Hand) *Message {
	return NewMessage(ShowHand{
		Pos:   pos,
		Cards: cards,
		Hand:  hand,
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

func NewWinner(player model.Player, amount float64) *Message {
	return NewMessage(Winner{
		Player: player,
		Amount: amount,
	})
}
