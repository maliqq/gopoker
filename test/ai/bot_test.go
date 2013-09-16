package ai

import (
	"fmt"
	"net/rpc/jsonrpc"
	"testing"
	"time"
)

import (
	"gopoker/model"
	"gopoker/model/game"
	"gopoker/server"
	rpc_service "gopoker/server/noderpc"
	"gopoker/util"
)

func TestBot(t *testing.T) {
	rpcAddr := "localhost:8081"
	zmqAddr := "tcp://127.0.0.1:5555"

	node := server.NewNode("bots_test", &server.Config{
		RPC: &server.RPCConfig{
			Addr: rpcAddr,
		},
		ZMQ: zmqAddr,
	})
	roomID := util.RandomUuid()
	tableSize := 9

	go node.StartRPC()
	go node.StartZMQ()

	<-time.After(1 * time.Second)

	client, err := jsonrpc.Dial("tcp", rpcAddr)
	if err != nil {
		t.Fatal("dialing error: ", err)
	}

	args := &rpc_service.CreateRoom{
		ID:        roomID,
		TableSize: tableSize,
		BetSize:   10.,
		Game:      model.NewGame(game.Texas, game.FixedLimit),
	}

	var result rpc_service.CallResult

	err = client.Call("NodeRPC.CreateRoom", args, &result)
	if err != nil {
		t.Fatal("call error: ", err)
	}
	t.Logf("rooms=%s", node.Rooms)

	bots := make([]*Bot, tableSize)
	for i := 0; i < tableSize; i++ {
		bot := NewBot(fmt.Sprintf("bot-%d", i), rpcAddr, zmqAddr)
		bots[i] = bot
	}
	t.Logf("bots=%#v", bots)
}
