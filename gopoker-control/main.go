package main

import (
	"flag"
	"log"
	"net/rpc"
)

import (
	"gopoker/model"
	"gopoker/model/game"
	"gopoker/server/service"
	"gopoker/util"
)

var (
	betSize     = flag.Float64("betsize", 20., "Bet size")
	limit       = flag.String("limit", "fixed-limit", "Limit to play")
	limitedGame = flag.String("game", "texas", "Game to play")
	mixedGame   = flag.String("mix", "", "Mix to play")
)

func main() {
	flag.Parse()

	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		log.Fatal("dialing error: ", err)
	}

	args := &service.CreateRoom{
		Id:        model.Id(util.RandomUuid()),
		TableSize: 9,
		BetSize:   *betSize,
	}

	if *mixedGame != "" {
		args.Mix = model.NewMix(game.MixedGame(*mixedGame))
	} else {
		args.Game = model.NewGame(game.LimitedGame(*limitedGame), game.Limit(*limit))
	}

	var result service.CallResult

	err = client.Call("NodeRPC.CreateRoom", args, &result)
	if err != nil {
		log.Fatal("call error: ", err)
	}
}
