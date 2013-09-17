package message

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/deal"
	"gopoker/poker"
)

// Event Message
type Message struct {
	Event     string
	Timestamp int64
	Payload   interface{}
}

func (msg *Message) Reset() {

}

func (msg *Message) ProtoMessage() {

}

func (msg *Message) String() string {
	return ""
}

type Format interface {
	/*
	NewAddBet(pos int, bet *Bet)
	NotifyRequireBet(pos int, betRange *bet.Range)
	NotifyBettingComplete(total float64)

	NotifyDealPocket(pos int, cards Cards, dealingType deal.Type)
	NotifyDealShared(cards Cards, dealingType deal.Type)
	NotifyRequireDiscard(req *RequireDiscard)
	NotifyDiscarded(pos int, cardsNum int)
	NotifyDiscardCards(pos int, cards Cards)

	NotifyPlayStart(play *Play)
	NotifyStreetStart(name string)
	NotifyPlayStop()

	NotifyShowHand(pos int, player *string, cards poker.Cards, hand *Hand, handStr string)
	NotifyShowCards(pos int, player *string, cards poker.Cards)
	NotifyMuckCards(pos int, player *string, cards poker.Cards)
	NotifyWinner(pos int, player *string, amount float64)

	NotifyMoveButton(pos int)
	NotifyJoinTable(player string, pos int, amount float64)
	NotifyLeaveTable(player string)

	NotifyErrorMessage(err error)
	NotifyChatMessage(body string)
	*/
}

type AddBet struct {
	Pos int
	Bet *model.Bet
}

type RequireBet struct {
	Pos   int
	Range *bet.Range
}

type BettingComplete struct {
	Total float64
}

type DealCards struct {
	Pos   int
	Cards poker.Cards
	Type  deal.Type
}

type RequireDiscard struct {
	Pos int
}

type Discarded struct {
	Pos int
	Num int
}

type Discard struct {
	Pos   int
	Cards poker.Cards
}

type JoinTable struct {
	Player model.Player
	Pos int
	Amount float64
}
