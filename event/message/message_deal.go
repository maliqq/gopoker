package message

import (
	"gopoker/model/deal"
	"gopoker/poker"
)

// DealCards
type DealCards struct {
	Pos   int
	Cards poker.Cards
	Type  deal.Type
}

func (DealCards) EventMessage() {}

// RequireDiscard
type RequireDiscard struct {
	Pos int
}

func (RequireDiscard) EventMessage() {}

// Discarded
type Discarded struct {
	Pos int
	Num int
}

func (Discarded) EventMessage() {}

// DiscardCards
type DiscardCards struct {
	Pos   int
	Cards poker.Cards
}

func (DiscardCards) EventMessage() {}
