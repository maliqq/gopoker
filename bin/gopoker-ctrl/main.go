package main

//
// create new play via JSONRPC
//
import (
	"flag"
	"math/rand"
	"time"
)

import (
	"gopoker/model"
	"gopoker/model/game"
	"gopoker/client/rpc_client"
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

	guid := model.Guid(*roomID)

	client := rpc_client.NewConnection("localhost:8081")

	var variation model.Variation
	if *mixedGame != "" {
		variation = model.NewMix(game.MixedGame(*mixedGame), *tableSize)
	} else {
		variation = model.NewGame(game.LimitedGame(*limitedGame), game.Limit(*limit), *tableSize)
	}

	client.CreateRoom(guid, *betSize, variation)
}
