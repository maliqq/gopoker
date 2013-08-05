package ai

import (
	"fmt"
	"log"
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

// Bot - bot
type Bot struct {
	ID      string
	roomID  string
	pos     int
	players []model.Player
	game    *model.Game
	street  string
	bet     float64
	pot     float64
	cards   poker.Cards
	board   poker.Cards
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
		cards:   poker.Cards{},
		board:   poker.Cards{},
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
	b.addBet(model.NewFold())
}

func (b *Bot) raise(amount float64) {
	b.addBet(model.NewRaise(amount))
}

func (b *Bot) call(amount float64) {
	b.addBet(model.NewCall(amount))
}

func (b *Bot) addBet(newBet *model.Bet) {
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

func (b *Bot) decide(betRange *message.BetRange) {
	switch b.street {
	case "preflop":
		b.decidePreflop(b.cards)
	case "flop", "turn", "river":
		b.decideBoard(b.cards, b.board)
	}
}

func (b *Bot) decidePreflop(cards poker.Cards) {
	group := calc.SklanskyMalmuthGroup(cards[0], cards[1])

	log.Printf("group=%d", group)

	switch group {
	case 9:
		// fold
	case 7, 8:
		// call BB
	case 5, 6:
		// raise BB..BB*4
		// raiseChance = 0.2
	case 3, 4:
		// raise stack+bet
		// raiseChance = 0.5
		// allInChance = 0.1
	case 1, 2:
		// raise stack+bet
		// raiseChance = 0.5
		// allInChance = 0.1
	}
}

func (b *Bot) decideBoard(cards, board poker.Cards) {
	opponentsNum := len(b.players) - 1
	chances := calc.ChancesAgainstN{OpponentsNum: opponentsNum}.WithBoard(cards, board)

	tightness := 0.
	if chances.Wins() > tightness {
		// raise stack+bet
		// raiseChance = 0.5
		// allInChance = 0.5
	} else if chances.Wins() > tightness / 2 {
		// raise (stack+bet)/3
		// raiseChance = 0.2
	} else if(chances.Ties() > 0.8) {
		// raise stack+bet
		// raiseChance = 0
		// allInChance = 0
  } else {
    // check/fold
  }
}
