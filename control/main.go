package main

import (
	"log"
	"net/rpc"
	"flag"
)

import (
	"gopoker/util"
	"gopoker/model"
	"gopoker/model/game"
	"gopoker/server/service"
)

var (
	stackSize   = flag.Float64("stacksize", 20., "Stack size")
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

	stake := game.NewStake(*stackSize)

	args := &service.CreateRoom{
		Id: model.Id(util.RandomUuid()),
		Size: 9,
	}

	if *mixedGame != "" {
		args.Mix = model.NewMix(game.MixedGame(*mixedGame), stake)
	} else {
		args.Game = model.NewGame(game.LimitedGame(*limitedGame), game.Limit(*limit), stake)
	}

	var result service.CallResult

	err = client.Call("NodeRPC.CreateRoom", args, &result)
	if err != nil {
		log.Fatal("call error: ", err)
	}
}
