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
	ID     string
	roomID string
	pos    int
	stack  float64

	*context

	client  *rpc.Client
	zmqConn *zeromq_client.Connection
}

// NewBot - create new bot
func NewBot(id, rpcAddr, sockAddr string) *Bot {
	if id == "" {
		id = util.RandomUuid()
	}
	log.SetPrefix(fmt.Sprintf("[bot#%s]", id))

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
	b.stack = amount

	log.Printf("joining table...")
	b.notifyRoom(message.NotifyJoinTable(b.ID, pos, amount))

	log.Printf("connecting gateway...")
	b.callRPC("ConnectGateway", rpc_service.ConnectGateway{
		RoomID:   roomID,
		PlayerID: b.ID,
	})
}

// Play - start bot
func (b *Bot) Play() {
	for msg := range b.zmqConn.Recv {
		//log.Printf("received msg: %s", msg)

		switch msg.Payload().(type) {
		case *message.PlayStart:
			start := msg.Envelope.PlayStart

			b.cards = poker.Cards{}
			b.board = poker.Cards{}
			b.pot = 0.
			b.opponentsNum = 3
			b.stake = model.NewStake(start.Play.Stake.GetBigBlind())

		case *message.StreetStart:
			b.street = msg.Envelope.StreetStart.GetName()

		case *message.BettingComplete:
			b.pot = msg.Envelope.BettingComplete.GetPot()
			b.bet = 0.

		case *message.DealCards:
			deal := msg.Envelope.DealCards
			switch deal.GetType() {
			case message.DealType_Board:
				b.board = b.board.Append(poker.BinaryCards(deal.Cards))

			case message.DealType_Hole:
				b.cards = b.cards.Append(poker.BinaryCards(deal.Cards))
			}

		case *message.RequireBet:
			req := msg.Envelope.RequireBet
			if int(req.GetPos()) == b.pos {
				// pause
				<-time.After(1 * time.Second)
				// our turn
				b.decide(req.BetRange)
			}

		case *message.AddBet:
			bet := msg.Envelope.AddBet
			if int(bet.GetPos()) == b.pos {
				b.bet = bet.Bet.GetAmount()
			}
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
	msg := message.NotifyAddBet(b.pos, newBet.Proto())
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

func (b *Bot) decidePreflop(cards poker.Cards) decision {
	group := calc.SklanskyMalmuthGroup(cards[0], cards[1])

	log.Printf("group=%d", group)

	bb := b.stake.BigBlindAmount()
	switch group {
	case 9:
		return decision{maxBet: 0.}

	case 7, 8:
		return decision{maxBet: bb}

	case 5, 6:
		return decision{
			minBet:      bb,
			maxBet:      bb * 4,
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
	chances := calc.ChancesAgainstN{OpponentsNum: b.opponentsNum}.WithBoard(cards, board)

	log.Printf("chances=%s", chances)

	tightness := 0.5
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

type action int

const (
	fold action = iota
	checkFold
	checkCall
	raise
	allIn
)

func (b *Bot) invoke(decision decision, betRange *message.BetRange) {
	call, minRaise, maxRaise := betRange.GetCall(), betRange.GetMin(), betRange.GetMax()

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
