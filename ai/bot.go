package ai

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

import (
	"gopoker/client/zmq_client"
	"gopoker/message"
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/deal"
	"gopoker/poker"
	"gopoker/util"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

type context struct {
	opponentsNum int
	game         *model.Game
	stake        *model.Stake
	street       string
	bet          float64
	pot          float64
	cards        poker.Cards
	board        poker.Cards
}

// Bot - bot
type Bot struct {
	ID     model.Player
	roomID string
	pos    int
	stack  float64

	*context

	zmqConn *zmq_client.NodeZMQ
}

// NewBot - create new bot
func NewBot(id, room, publisher, receiver string) *Bot {
	if id == "" {
		id = util.RandomUuid()
	}
	log.SetPrefix(fmt.Sprintf("[bot#%s] ", id))

	return &Bot{
		ID:      model.Player(id),
		roomID:  room,
		zmqConn: zmq_client.ConnectZmqGateway(publisher, receiver, id),
		context: &context{
			cards: poker.Cards{},
			board: poker.Cards{},
		},
	}
}

// Join - join player
func (b *Bot) Join(pos int, amount float64) {
	b.pos = pos
	b.stack = amount

	log.Printf("joining table...")

	b.sendMultipart(&message.Join{b.ID, pos, amount})
}

// Play - start bot
func (b *Bot) Play() {
	for data := range b.zmqConn.Recv {
		event := b.receiveMultipart(data)
		log.Printf("received: %s", event)

		switch msg := event.Message.(type) {
		case *message.PlayStart:

			b.cards = poker.Cards{}
			b.board = poker.Cards{}
			b.pot = 0.
			b.opponentsNum = 6
			b.stake = msg.Stake

		case *message.StreetStart:

			b.street = msg.Name

		case *message.BettingComplete:

			b.pot = msg.Pot
			b.bet = 0.

		case *message.DealCards:

			switch msg.Type {
			case deal.Board:
				b.board = b.board.Append(msg.Cards)

			case deal.Hole:
				b.cards = b.cards.Append(msg.Cards)
			}

		case *message.RequireBet:

			if msg.Pos == b.pos {
				// pause
				<-time.After(1 * time.Second)
				// our turn
				b.decide(msg.Range)
			}

		case *message.AddBet:

			if msg.Pos == b.pos {
				b.bet = msg.Bet.Amount
			}
		}
	}
}

func (b *Bot) sendMultipart(msg message.Message) {
	data, err := proto.Marshal(event.New(msg).Proto())
	if err != nil {
		log.Printf("marshal error: %s", err)
	} else {
		multipart := [][]byte{
			[]byte(b.ID),
			[]byte(b.roomID),
			data,
		}
		//log.Printf("sending %d bytes", len(data))
		b.zmqConn.Send <- multipart
	}
}

func (b *Bot) receiveMultipart(multipart [][]byte) *event.Event {
	//topic := multipart[0]
	data := multipart[1]
	//log.Printf("received %d bytes for %s", len(data), topic)

	event := &event.Event{}
	if err := event.Unproto(data); err != nil {
		log.Printf("unmarshal error: %s", err)
		return nil
	}

	return event
}

func (b *Bot) check() {
	b.addBet(model.NewCheck())
}

func (b *Bot) fold() {
	b.bet = 0.
	b.addBet(model.NewFold())
}

func (b *Bot) raise(amount float64) {
	b.stack = b.stack + b.bet - amount
	b.bet = amount
	b.addBet(model.NewRaise(amount))
}

func (b *Bot) call(amount float64) {
	b.stack = b.stack + b.bet - amount
	b.bet = amount
	b.addBet(model.NewCall(amount))
}

func (b *Bot) addBet(newBet *model.Bet) {
	log.Printf("=== %s", newBet)
	b.sendMultipart(&message.AddBet{b.pos, newBet})
}

func (b *Bot) decide(betRange *bet.Range) {
	var decision decision

	if len(b.cards) != 2 {
		log.Printf("*** can't decide with cards=%s", b.cards)
		b.fold()
		return
	}

	if len(b.board) == 0 {
		decision = b.decidePreflop(b.cards)
	} else {
		decision = b.decideBoard(b.cards, b.board)
	}

	b.invoke(decision, betRange)
}

type action int

const (
	fold action = iota
	checkFold
	checkCall
	raise
	allIn
)

func (b *Bot) invoke(decision decision, betRange *bet.Range) {
	call, minRaise, maxRaise := betRange.Call, betRange.Min, betRange.Max

	log.Printf("decision=%#v call=%.2f minRaise=%.2f maxRaise=%.2f", decision, call, minRaise, maxRaise)

	min := call
	if min > b.stack+b.bet {
		min = b.stack + b.bet
	}
	max := decision.maxBet

	var action action
	if minRaise == 0. && maxRaise == 0. {
		action = checkCall
	} else if min > max {
		action = checkFold
	} else {
		if rand.Float64() < decision.raiseChance {
			action = raise
		} else if rand.Float64() < decision.allInChance {
			action = allIn
		} else {
			action = checkCall
		}
	}

	switch action {
	case fold:
		b.fold()

	case checkFold:
		if call == b.bet {
			b.check()
		} else {
			b.fold()
		}

	case checkCall:
		if call == b.bet || call == 0. {
			b.check()
		} else if call > 0. {
			b.call(call)
		}

	case raise:
		if minRaise == maxRaise {
			// raise fixed limit
			b.raise(maxRaise)
		} else {
			// raise no limit/pot limit
			d := maxRaise - minRaise
			bb := b.stake.BigBlindAmount()
			amount := minRaise + d*rand.Float64()
			amount = math.Floor(amount/bb) * bb // FIXME
			b.raise(amount)
		}

	case allIn:
		// all in
		if minRaise == maxRaise {
			// raise fixed limit
			b.raise(maxRaise)
		} else {
			// raise fixed limit
			b.raise(b.stack + b.bet)
		}
	}
}
