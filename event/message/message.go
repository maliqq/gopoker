package message

import (
	"gopoker/event/message/protobuf"
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/deal"
	"gopoker/poker"
)

// Event Message
type Message interface {
	EventMessage()
	Proto() *protobuf.Message
}

// AddBet
type AddBet struct {
	Pos int
	Bet *model.Bet
}

func (AddBet) EventMessage() {}

// RequireBet
type RequireBet struct {
	Pos   int
	Range *bet.Range
}

func (RequireBet) EventMessage() {}

// BettingComplete
type BettingComplete struct {
	Pot float64
}

func (BettingComplete) EventMessage() {}

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

// PlayStart
type PlayStart struct {
	Game  *model.Game
	Stake *model.Stake
	Table *model.Table
}

func (PlayStart) EventMessage() {}

// StreetStart
type StreetStart struct {
	Name string
}

func (StreetStart) EventMessage() {}

// PlayStop
type PlayStop struct {
}

func (PlayStop) EventMessage() {}

// ShowHand
type ShowHand struct {
	Pos      int
	Player   model.Player
	Cards    poker.Cards
	Hand     poker.Hand
	HandName string
}

func (ShowHand) EventMessage() {}

// ShowCards
type ShowCards struct {
	Pos    int
	Muck   bool
	Cards  poker.Cards
	Player model.Player
}

func (ShowCards) EventMessage() {}

// Winner
type Winner struct {
	Pos    int
	Player model.Player
	Amount float64
}

func (Winner) EventMessage() {}

// MoveButton
type MoveButton struct {
	Pos int
}

func (MoveButton) EventMessage() {}

// JoinTable
type JoinTable struct {
	Player model.Player
	Pos    int
	Amount float64
}

func (JoinTable) EventMessage() {}

// SitOut
type SitOut struct {
	Pos int
}

func (SitOut) EventMessage() {}

// ComeBack
type ComeBack struct {
	Pos int
}

func (ComeBack) EventMessage() {}

// LeaveTable
type LeaveTable struct {
	Player model.Player
}

func (LeaveTable) EventMessage() {}

// ErrorMessage
type ErrorMessage struct {
	Error error
}

func (ErrorMessage) EventMessage() {}

// ChatMessage
type ChatMessage struct {
	Body string
}

func (ChatMessage) EventMessage() {}
