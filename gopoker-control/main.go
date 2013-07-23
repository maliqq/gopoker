package main

import (
	"fmt"
	"flag"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"math/rand"
)

import (
	"gopoker/model"
	"gopoker/model/game"
	"gopoker/server/rpc_service"
	"gopoker/protocol"
	_"gopoker/util"
)

var (
	tableSize   = flag.Int("tablesize", 9, "Table size")
	betSize     = flag.Float64("betsize", 20., "Bet size")
	limit       = flag.String("limit", "fixed-limit", "Limit to play")
	limitedGame = flag.String("game", "texas", "Game to play")
	mixedGame   = flag.String("mix", "", "Mix to play")
)

func main() {
	flag.Parse()

	client, err := jsonrpc.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal("dialing error: ", err)
	}

	roomId := "0"
	args := &rpc_service.CreateRoom{
		Id:        roomId,
		TableSize: *tableSize,
		BetSize:   *betSize,
	}

	if *mixedGame != "" {
		args.Mix = model.NewMix(game.MixedGame(*mixedGame))
	} else {
		args.Game = model.NewGame(game.LimitedGame(*limitedGame), game.Limit(*limit))
	}

	call(client, "NodeRPC.CreateRoom", args)

	for pos := 0; pos < *tableSize; pos++ {
		player := fmt.Sprintf("player-%d", pos)
		amount := float64(rand.Intn(1000) + 1000)
		call(client, "NodeRPC.NotifyRoom", &rpc_service.NotifyRoom{
			Id: roomId,
			Message: protocol.NewJoinTable(model.Player(player), pos, amount),
		})
	}

	call(client, "NodeRPC.StartRoom", &rpc_service.StartRoom{
		Id: roomId,
	})
}

func call(client *rpc.Client, method string, args interface{}) {
	var result rpc_service.CallResult

	err := client.Call(method, args, &result)
	if err != nil {
		log.Fatal("call error: ", err)
	}
}
