package ai

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

import (
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
	game    *model.Game
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
	}
}

func (b *Bot) Join(roomID string, pos int, amount float64) {
	b.roomID = roomID
	b.pos = pos

	log.Printf("joining table...")
	b.notifyRoom(message.NewJoinTable(b.ID, pos, amount))

	log.Printf("connecting gateway...")
	b.call("ConnectGateway", rpc_service.ConnectGateway{
		RoomID:   roomID,
		PlayerID: b.ID,
	})
}

func (b *Bot) Play() {
	for msg := range b.zmqConn.Recv {
		log.Printf("received msg: %s", msg)

		switch msg.Payload().(type) {
		case *message.RequireBet:
		case *message.AddBet:
		}
	}
}

func (b *Bot) Check() {
	b.addBet(model.NewCheck())
}

func (b *Bot) Fold() {
	b.addBet(model.NewFold())
}

func (b *Bot) Raise(amount float64) {
	b.addBet(model.NewRaise(amount))
}

func (b *Bot) Call(amount float64) {
	b.addBet(model.NewCall(amount))
}

func (b *Bot) addBet(newBet *model.Bet) {
	msg := message.NewAddBet(b.pos, newBet.Proto())
	b.notifyRoom(msg)
}

func (b *Bot) notifyRoom(msg *message.Message) {
	b.call("NotifyRoom", rpc_service.NotifyRoom{
		ID:      b.roomID,
		Message: msg,
	})
}

func (b *Bot) call(method string, args interface{}) {
	var result rpc_service.CallResult

	err := b.client.Call(fmt.Sprintf("NodeRPC.%s", method), args, &result)

	if err != nil {
		log.Fatal("rpc call error: ", err)
	}
}
