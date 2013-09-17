package message

import (
	"gopoker/event/message/format/protobuf"
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
func (m AddBet) Proto() *protobuf.Message {
	return protobuf.NewAddBet(m.Pos, m.Bet.Proto())
}

// RequireBet
type RequireBet struct {
	Pos   int
	Range *bet.Range
}

func (RequireBet) EventMessage() {}
func (m RequireBet) Proto() *protobuf.Message {
	return protobuf.NewRequireBet(m.Pos, m.Range.Proto())
}

// BettingComplete
type BettingComplete struct {
	Pot float64
}

func (BettingComplete) EventMessage() {}
func (m BettingComplete) Proto() *protobuf.Message {
	return protobuf.NewBettingComplete(m.Pot)
}

// DealCards
type DealCards struct {
	Pos   int
	Cards poker.Cards
	Type  deal.Type
}

func (DealCards) EventMessage() {}
func (m DealCards) Proto() *protobuf.Message {
	return protobuf.NewDealCards(m.Pos, m.Cards.Proto(), m.Type)
}

// RequireDiscard
type RequireDiscard struct {
	Pos int
}

func (RequireDiscard) EventMessage() {}
func (m RequireDiscard) Proto() *protobuf.Message {
	return protobuf.NewRequireDiscard(m.Pos)
}

// Discarded
type Discarded struct {
	Pos int
	Num int
}

func (Discarded) EventMessage() {}
func (m Discarded) Proto() *protobuf.Message {
	return protobuf.NewDiscarded(m.Pos, m.Num)
}

// DiscardCards
type DiscardCards struct {
	Pos   int
	Cards poker.Cards
}

func (DiscardCards) EventMessage() {}
func (m DiscardCards) Proto() *protobuf.Message {
	return protobuf.NewDiscardCards(m.Pos, m.Cards.Proto())
}

// PlayStart
type PlayStart struct {
	Game  *model.Game
	Stake *model.Stake
	Table *model.Table
}

func (PlayStart) EventMessage() {}
func (m PlayStart) Proto() *protobuf.Message {
	return protobuf.NewPlayStart(nil)
}

// StreetStart
type StreetStart struct {
	Name string
}

func (StreetStart) EventMessage() {}
func (m StreetStart) Proto() *protobuf.Message {
	return protobuf.NewStreetStart(m.Name)
}

// PlayStop
type PlayStop struct {
}

func (PlayStop) EventMessage() {}
func (m PlayStop) Proto() *protobuf.Message {
	return protobuf.NewPlayStop() // FIXME
}

// ShowHand
type ShowHand struct {
	Pos      int
	Player   model.Player
	Cards    poker.Cards
	Hand     poker.Hand
	HandName string
}

func (ShowHand) EventMessage() {}
func (m ShowHand) Proto() *protobuf.Message {
	return protobuf.NewShowHand(
		m.Pos,
		m.Player.Proto(),
		m.Cards.Proto(),
		m.Hand.Proto(),
		m.HandName,
	)
}

// ShowCards
type ShowCards struct {
	Pos    int
	Muck   bool
	Cards  poker.Cards
	Player model.Player
}

func (ShowCards) EventMessage() {}
func (m ShowCards) Proto() *protobuf.Message {
	return protobuf.NewShowCards(m.Pos, m.Player.Proto(), m.Cards.Proto())
}

// Winner
type Winner struct {
	Pos    int
	Player model.Player
	Amount float64
}

func (Winner) EventMessage() {}
func (m Winner) Proto() *protobuf.Message {
	return protobuf.NewWinner(m.Pos, m.Player.Proto(), m.Amount)
}

// MoveButton
type MoveButton struct {
	Pos int
}

func (MoveButton) EventMessage() {}
func (m MoveButton) Proto() *protobuf.Message {
	return protobuf.NewMoveButton(m.Pos)
}

// JoinTable
type JoinTable struct {
	Player model.Player
	Pos    int
	Amount float64
}

func (JoinTable) EventMessage() {}
func (m JoinTable) Proto() *protobuf.Message {
	return protobuf.NewJoinTable(m.Player.Proto(), m.Pos, m.Amount)
}

// SitOut
type SitOut struct {
	Pos int
}

func (SitOut) EventMessage() {}
func (m SitOut) Proto() *protobuf.Message {
	return protobuf.NewSitOut(m.Pos)
}

// ComeBack
type ComeBack struct {
	Pos int
}

func (ComeBack) EventMessage() {}
func (m ComeBack) Proto() *protobuf.Message {
	return protobuf.NewComeBack(m.Pos)
}

// LeaveTable
type LeaveTable struct {
	Player model.Player
}

func (LeaveTable) EventMessage() {}
func (m LeaveTable) Proto() *protobuf.Message {
	return protobuf.NewLeaveTable(m.Player.Proto())
}

// ErrorMessage
type ErrorMessage struct {
	Error error
}

func (ErrorMessage) EventMessage() {}
func (m ErrorMessage) Proto() *protobuf.Message {
	return protobuf.NewErrorMessage(m.Error)
}

// ChatMessage
type ChatMessage struct {
	Body string
}

func (ChatMessage) EventMessage() {}
func (m ChatMessage) Proto() *protobuf.Message {
	return protobuf.NewChatMessage(m.Body)
}
