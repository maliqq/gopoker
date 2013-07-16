package protocol

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

import (
	"gopoker/model"
	"gopoker/model/deal"
	"gopoker/model/game"
	"gopoker/model/seat"
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

const (
	UseIndent = false
)

func (msg *Message) String() string {
	var err error
	var s []byte
	if UseIndent {
		s, err = json.MarshalIndent(msg, "", "\t")
	} else {
		s, err = json.Marshal(msg)
	}

	if err != nil {
		log.Printf("Message: %#v\n", msg)
		log.Printf("Error: %s\n", err)

		panic("error marshaling message")
	}

	return string(s)
}

// error
type Error struct {
	Description string
}

func NewError(err error) *Message {
	return NewMessage(
		Error{
			Description: err.Error(),
		},
	)
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
	Player *model.Player
	Pos    int
	Amount float64
}

func NewJoinTable(player *model.Player, pos int, amount float64) *Message {
	return NewMessage(
		JoinTable{
			Player: player,
			Pos:    pos,
			Amount: amount,
		},
	)
}

type LeaveTable struct {
	Player *model.Player
}

func NewLeaveTable(player *model.Player) *Message {
	return NewMessage(
		LeaveTable{
			Player: player,
		},
	)
}

type ChangeTableState struct {
	State string
}

type SeatState struct {
	Pos int
}

type ChangeSeatState struct {
	Pos   int
	State seat.State
}

type SitOut struct {
	Pos int
}

type ComeBack struct {
	Pos int
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

func NewRequireDiscard(req *RequireDiscard) *Message {
	return NewMessage(*req)
}

type Discarded struct {
	Pos int
	Num int
}

func NewDiscarded(pos int, cardsNum int) *Message {
	return NewMessage(
		Discarded{
			Pos: pos,
			Num: cardsNum,
		},
	)
}

type DiscardCards struct {
	Pos   int
	Cards poker.Cards
}

func NewDiscardCards(pos int, cards *poker.Cards) *Message {
	return NewMessage(
		DiscardCards{
			Pos:   pos,
			Cards: *cards,
		},
	)
}

type RequireBet struct {
	Pos int
	model.BetRange
}

func (r RequireBet) String() string {
	return fmt.Sprintf("call: %.2f min: %.2f max: %.2f", r.Call, r.Min, r.Max)
}

type AddBet struct {
	Pos int
	Bet model.Bet
}

func NewAddBet(pos int, bet *model.Bet) *Message {
	return NewMessage(
		AddBet{
			Pos: pos,
			Bet: *bet,
		},
	)
}

type PotSummary struct {
	Amount float64
}

func NewPotSummary(pot *model.Pot) *Message {
	return NewMessage(
		PotSummary{
			Amount: pot.Total(),
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
	Type game.LimitedGame
	game.Limit
}

func NewChangeGame(g *model.Game) *Message {
	return NewMessage(
		ChangeGame{
			Type:  g.Type,
			Limit: g.Limit,
		},
	)
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
