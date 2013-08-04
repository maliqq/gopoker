package main

//
// create new play via JSONRPC
//
import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

import (
	"gopoker/model"
	"gopoker/model/game"
	"gopoker/protocol/message"
	rpc_service "gopoker/server/noderpc"
	_ "gopoker/util"
)

var (
	tableSize   = flag.Int("tablesize", 9, "Table size")
	betSize     = flag.Float64("betsize", 20., "Bet size")
	limit       = flag.String("limit", "FixedLimit", "Limit to play")
	limitedGame = flag.String("game", "Texas", "Game to play")
	mixedGame   = flag.String("mix", "", "Mix to play")
	roomID      = flag.String("roomid", "0", "Set Room ID")
)

const (
	defaultConfigDir = "/etc/gopoker"
)

var (
	configDir = flag.String("config-dir", defaultConfigDir, "Config dir")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	model.LoadGames(*configDir)

	client, err := jsonrpc.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal("dialing error: ", err)
	}

	args := &rpc_service.CreateRoom{
		ID:        *roomID,
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
			ID:      *roomID,
			Message: message.NewJoinTable(player, pos, amount),
		})
	}

	call(client, "NodeRPC.StartRoom", &rpc_service.StartRoom{
		ID: *roomID,
	})
}

func call(client *rpc.Client, method string, args interface{}) {
	var result rpc_service.CallResult

	err := client.Call(method, args, &result)
	if err != nil {
		log.Fatal("call error: ", err)
	}
}
