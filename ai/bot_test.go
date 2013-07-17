package ai

import (
	"testing"
	"net/rpc"
)

import (
	"gopoker/model"
	"gopoker/model/game"
	"gopoker/server"
	"gopoker/util"
	"gopoker/server/service"
)

func TestBot(t *testing.T) {
	rpcAddr := "localhost:8081"
	node := server.NewNode("bots_test", "", rpcAddr)
	roomId := util.RandomUuid()
	tableSize := 9

	node.StartRPC()

	client, err := rpc.DialHTTP("tcp", rpcAddr)
	if err != nil {
		t.Fatalf("dialing error: ", err)
	}

	args := &service.CreateRoom{
		Id:        model.Id(roomId),
		TableSize: tableSize,
		BetSize:   10.,
		Game:      model.NewGame(game.Texas, game.FixedLimit),
	}

	var result service.CallResult

	err = client.Call("NodeRPC.CreateRoom", args, &result)
	if err != nil {
		t.Fatalf("call error: ", err)
	}
	t.Logf("rooms=%s", node.Rooms)

	bots := make([]*Bot, tableSize)
	for i := 0; i < tableSize; i++ {
		bot := NewBot(rpcAddr)
		bots[i] = bot
		bot.Join(roomId, i)
	}
	t.Logf("bots=%#v", bots)
}
