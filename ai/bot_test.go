package ai

import (
	"net/rpc"
	"testing"
)

import (
	"gopoker/model"
	"gopoker/model/game"
	"gopoker/server"
	"gopoker/server/rpc_service"
	"gopoker/util"
)

func TestBot(t *testing.T) {
	rpcAddr := "localhost:8081"
	node := server.NewNode("bots_test", "", rpcAddr)
	roomID := util.RandomUuid()
	tableSize := 9

	node.StartRPC()

	client, err := rpc.DialHTTP("tcp", rpcAddr)
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
		bot := NewBot(rpcAddr)
		bots[i] = bot
	}
	t.Logf("bots=%#v", bots)
}
