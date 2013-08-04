package ai

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

import (
	zeromq_client "gopoker/client/zmq"
	"gopoker/protocol/message"
	rpc_service "gopoker/server/noderpc"
	"gopoker/util"
)

// Bot - bot
type Bot struct {
	ID      string
	client  *rpc.Client
	zmqConn *zeromq_client.Connection
}

// NewBot - create new bot
func NewBot(rpcAddr, sockAddr string) *Bot {
	id := util.RandomUuid()

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
	var result rpc_service.CallResult

	err := b.client.Call("NodeRPC.NotifyRoom", rpc_service.NotifyRoom{
		ID:      roomID,
		Message: message.NewJoinTable(b.ID, pos, amount),
	}, &result)

	if err != nil {
		log.Fatal("rpc error: ", err)
	}
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
