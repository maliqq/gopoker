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
	"gopoker/event"
	"gopoker/event/message"
	"gopoker/model"
	"gopoker/model/game"
	rpc_service "gopoker/server/noderpc"
	_ "gopoker/util"
)

var (
	tableSize     = flag.Int("tablesize", 9, "Table size")
	betSize       = flag.Float64("betsize", 20., "Bet size")
	limit         = flag.String("limit", "FixedLimit", "Limit to play")
	limitedGame   = flag.String("game", "Texas", "Game to play")
	mixedGame     = flag.String("mix", "", "Mix to play")
	roomID        = flag.String("roomid", "0", "Set Room ID")
	createPlayers = flag.Bool("create", false, "Create players")
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

	guid := model.Guid(*roomID)

	client, err := jsonrpc.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal("dialing error: ", err)
	}

	args := &rpc_service.CreateRoom{
		Guid:    guid,
		BetSize: *betSize,
	}

	if *mixedGame != "" {
		args.Mix = model.NewMix(game.MixedGame(*mixedGame), *tableSize)
	} else {
		args.Game = model.NewGame(game.LimitedGame(*limitedGame), game.Limit(*limit), *tableSize)
	}

	call(client, "NodeRPC.CreateRoom", args)

	if *createPlayers {
		for pos := 0; pos < *tableSize; pos++ {
			player := fmt.Sprintf("player-%d", pos)
			amount := float64(rand.Intn(1000) + 1000)
			call(client, "NodeRPC.NotifyRoom", &rpc_service.NotifyRoom{
				Guid:  model.Guid(*roomID),
				Event: event.NewEvent(&message.JoinTable{model.Player(player), pos, amount}),
			})
		}
	}

	call(client, "NodeRPC.StartRoom", &rpc_service.StartRoom{
		Guid: guid,
	})
}

func call(client *rpc.Client, method string, args interface{}) {
	var result rpc_service.CallResult

	err := client.Call(method, args, &result)
	if err != nil {
		log.Fatal("call error: ", err)
	}
}
