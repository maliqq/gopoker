package ai

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

import (
	"gopoker/calc"
	zeromq_client "gopoker/client/zmq"
	"gopoker/model"
	"gopoker/poker"
	"gopoker/protocol/message"
	rpc_service "gopoker/server/noderpc"
	"gopoker/util"
)

type context struct {
	players []model.Player
	game    *model.Game
	street  string
	bb      float64
	stack   float64
	bet     float64
	pot     float64
	cards   poker.Cards
	board   poker.Cards
}

// Bot - bot
type Bot struct {
	ID     string
	roomID string
	pos    int

	*context

	client  *rpc.Client
	zmqConn *zeromq_client.Connection
}

// NewBot - create new bot
func NewBot(id, rpcAddr, sockAddr string) *Bot {
	if id == "" {
		id = util.RandomUuid()
	}

	client, err := jsonrpc.Dial("tcp", rpcAddr)
	if err != nil {
		log.Fatal("dialing error: ", err)
	}

	return &Bot{
		ID:      id,
		client:  client,
		zmqConn: zeromq_client.NewConnection(sockAddr, id),
		context: &context{
			cards: poker.Cards{},
			board: poker.Cards{},
		},
	}
}

// Join - join player
func (b *Bot) Join(roomID string, pos int, amount float64) {
	b.roomID = roomID
	b.pos = pos

	log.Printf("joining table...")
	b.notifyRoom(message.NewJoinTable(b.ID, pos, amount))

	log.Printf("connecting gateway...")
	b.callRPC("ConnectGateway", rpc_service.ConnectGateway{
		RoomID:   roomID,
		PlayerID: b.ID,
	})
}

// Play - start bot
func (b *Bot) Play() {
	for msg := range b.zmqConn.Recv {
		log.Printf("received msg: %s", msg)

		switch msg.Payload().(type) {
		case message.PlayStart:
			b.cards = poker.Cards{}
			b.board = poker.Cards{}
			b.pot = 0.

		case message.StreetStart:
			b.street = msg.Envelope.StreetStart.GetName()

		case message.BettingComplete:
			b.pot = msg.Envelope.BettingComplete.GetPot()

		case *message.DealCards:
			deal := msg.Envelope.DealCards
			switch deal.GetType() {
			case message.DealType_Board:
				b.board = append(b.board, poker.BinaryCards(deal.Cards)...)
			case message.DealType_Hole:
				b.cards = append(b.cards, poker.BinaryCards(deal.Cards)...)
			}

		case *message.RequireBet:
			req := msg.Envelope.RequireBet
			if int(req.GetPos()) == b.pos {
				// pause
				<-time.After(2 * time.Second)
				// our turn
				b.decide(req.BetRange)
			}

		case *message.AddBet:

		default:
			log.Printf("got %s", msg)
		}
	}
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
	msg := message.NewAddBet(b.pos, newBet.Proto())
	b.notifyRoom(msg)
}

func (b *Bot) notifyRoom(msg *message.Message) {
	b.callRPC("NotifyRoom", rpc_service.NotifyRoom{
		ID:      b.roomID,
		Message: msg,
	})
}

func (b *Bot) callRPC(method string, args interface{}) {
	var result rpc_service.CallResult

	err := b.client.Call(fmt.Sprintf("NodeRPC.%s", method), args, &result)

	if err != nil {
		log.Fatal("rpc call error: ", err)
	}
}

type decision struct {
	minBet      float64
	maxBet      float64
	raiseChance float64
	allInChance float64
}

func (b *Bot) decide(betRange *message.BetRange) {
	var decision decision

	switch b.street {
	case "preflop":
		decision = b.decidePreflop(b.cards)

	case "flop", "turn", "river":
		b.decideBoard(b.cards, b.board)
	}

	b.invoke(decision, betRange)
}

func (b *Bot) decidePreflop(cards poker.Cards) decision {
	group := calc.SklanskyMalmuthGroup(cards[0], cards[1])

	log.Printf("group=%d", group)

	switch group {
	case 9:
		return decision{maxBet: 0.}

	case 7, 8:
		return decision{maxBet: b.bb}

	case 5, 6:
		return decision{
			minBet:      b.bb,
			maxBet:      b.bb * 4,
			raiseChance: 0.2,
		}

	case 3, 4:
		return decision{
			maxBet:      b.stack + b.bet,
			raiseChance: 0.5,
			allInChance: 0.1,
		}

	case 1, 2:
		return decision{
			maxBet:      b.stack + b.bet,
			raiseChance: 0.5,
			allInChance: 0.1,
		}
	}

	return decision{}
}

func (b *Bot) decideBoard(cards, board poker.Cards) decision {
	opponentsNum := len(b.players) - 1
	chances := calc.ChancesAgainstN{OpponentsNum: opponentsNum}.WithBoard(cards, board)

	tightness := 0.7
	if chances.Wins() > tightness {
		return decision{
			maxBet:      b.stack + b.bet,
			raiseChance: 0.5,
			allInChance: 0.5,
		}
	} else if chances.Wins() > tightness/2 {
		return decision{
			maxBet:      (b.stack + b.bet) / 3.,
			raiseChance: 0.2,
		}

	} else if chances.Ties() > 0.8 {
		return decision{
			maxBet:      b.stack + b.bet,
			raiseChance: 0.,
			allInChance: 0.,
		}
	}

	return decision{}
}

func (b *Bot) invoke(decision decision, betRange *message.BetRange) {
	call, minRaise, maxRaise := betRange.GetCall(), betRange.GetMin(), betRange.GetMax()

	min := call
	if min > b.stack+b.bet {
		min = b.stack + b.bet
	}
	max := decision.maxBet

	if min > max {
		// check/fold
		if call > 0. {
			b.fold()
		} else {
			b.check()
		}
	} else {
		if rand.Float64() < decision.raiseChance {
			if minRaise == maxRaise {
				// raise fixed limit
				b.raise(maxRaise)
			} else {
				// raise no limit/pot limit
				d := maxRaise - minRaise
				amount := minRaise + d*rand.Float64()
				amount = math.Floor(amount/b.bb) * b.bb
				b.raise(amount)
			}
		} else if rand.Float64() < decision.allInChance {
			// all in
			b.raise(b.stack + b.bet)
		} else {
			// call
			if call > 0. {
				b.call(call)
			} else {
				b.check()
			}
		}
	}
}
