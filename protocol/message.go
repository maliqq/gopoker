package protocol

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/deal"
	"gopoker/poker"
)

type Payload interface{}

type Message struct {
	Type      string
	Timestamp int64
	Payload   Payload
}

func NewMessage(payload Payload) *Message {
	name := reflect.TypeOf(payload).Name()

	if name == "" {
		fmt.Printf("payload: %#v", payload)
		panic("unknown message type")
	}

	return &Message{
		Type:      name,
		Timestamp: time.Now().Unix(),
		Payload:   payload,
	}
}

func (msg *Message) String() string {
	s, _ := json.Marshal(msg)

	return string(s)
}

// table info
type Table struct{}

type MoveButton struct {
	Pos int
}

func NewMoveButton(pos int) *Message {
	return NewMessage(
		MoveButton{
			Pos: pos,
		},
	)
}

type JoinTable struct {
	Pos    int
	Amount float64
}

type LeaveTable struct {
	Pos int
}

type ChangeTableState struct {
	State string
}

type SeatState struct {
	Pos int
}

type ChangeSeatState struct {
	Pos   int
	State string
}

// seat info
type Seat struct {
	State string
	Stack float64
	Bet   float64
}

type SeatStack struct {
	Pos int
}

type ChangeSeatStack struct {
	Pos    int
	Amount float64
}

type AdvanceSeatStack struct {
	Pos    int
	Amount float64
}

type DealCards struct {
	Pos   int
	Cards poker.Cards
	Type  deal.Type
}

func NewDealPocket(pos int, cards *poker.Cards, dealingType deal.Type) *Message {
	return NewMessage(
		DealCards{
			Pos:   pos,
			Cards: *cards,
			Type:  dealingType,
		},
	)
}

func NewDealShared(cards *poker.Cards, dealingType deal.Type) *Message {
	return NewMessage(
		DealCards{
			Cards: *cards,
			Type:  dealingType,
		},
	)
}

type RequireDiscard struct {
	Pos int
}

func NewRequireDiscard(pos int) *Message {
	return NewMessage(
		RequireDiscard{
			Pos: pos,
		},
	)
}

type DiscardCards struct {
	Pos int
	Num int
}

func NewDiscardCards(pos int, cardsNum int) *Message {
	return NewMessage(
		DiscardCards{
			Pos: pos,
			Num: cardsNum,
		},
	)
}

type RequireBet struct {
	Pos  int
	Call float64
	Min  float64
	Max  float64
}

func NewRequireBet(req *RequireBet) *Message {
	return NewMessage(*req)
}

func (r RequireBet) String() string {
	return fmt.Sprintf("call: %.2f min: %.2f max: %.2f", r.Call, r.Min, r.Max)
}

type AddBet struct {
	Pos int
	Bet bet.Bet
}

func NewAddBet(pos int, bet *bet.Bet) *Message {
	return NewMessage(
		AddBet{
			Pos: pos,
			Bet: *bet,
		},
	)
}

type Chat struct {
	Pos     int
	Message string
}

// deal info
type Deal struct {
}

type ChangeDealState struct {
	State string
}

type ChangeGame struct {
}

// new street
type Street struct {
}

// hand info
type ShowHand struct {
	Pos   int
	Cards poker.Cards
	Hand  *poker.Hand
}

func NewShowHand(pos int, cards *poker.Cards, hand *poker.Hand) *Message {
	return NewMessage(
		ShowHand{
			Pos:   pos,
			Cards: *cards,
			Hand:  hand,
		},
	)
}

// pocket cards show
type ShowCards struct {
	Pos   int
	Cards poker.Cards
	Muck  bool
}

func NewShowCards(pos int, cards *poker.Cards) *Message {
	return NewMessage(
		ShowCards{
			Pos:   pos,
			Cards: *cards,
		},
	)
}

func NewMuckCards(pos int, cards *poker.Cards) *Message {
	return NewMessage(
		ShowCards{
			Pos:   pos,
			Cards: *cards,
			Muck:  true,
		},
	)
}

// win info
type Winner struct {
	Player model.Id
	Amount float64
}

func NewWinner(player model.Id, amount float64) *Message {
	return NewMessage(
		Winner{
			Player: player,
			Amount: amount,
		},
	)
}
