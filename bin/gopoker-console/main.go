package main

//
// command line play
//
import (
	"flag"
	"fmt"
	"log"
	"os"
)

import (
	"gopoker/protocol"

	"gopoker/model"
	"gopoker/model/game"

	"gopoker/client/console_client"
	"gopoker/play"
)

var (
	logfile     = flag.String("logfile", "", "Log file path")
	betsize     = flag.Float64("betsize", 20., "Bet size")
	mixedGame   = flag.String("mix", "", "Mix to play")
	limitedGame = flag.String("game", "texas", "Game to play")
)

func main() {
	flag.Parse()

	if *logfile != "" {
		w, err := os.Create(*logfile)

		if err != nil {
			panic(err.Error())
		}

		defer w.Close()
		log.SetOutput(w)
	}

	me := make(protocol.MessageChannel)
	play := createPlay(me)
	fmt.Printf("%s\n", play)

	go play.Start()

	conn := &console_client.Connection{
		Server: play,
	}

	for msg := range me {
		conn.Handle(msg)
	}
}

func createPlay(me protocol.MessageChannel) *play.Play {
	size := 3
	table := model.NewTable(size)
	stake := model.NewStake(*betsize)
	//stake.WithAnte = true

	var variation model.Variation
	if *mixedGame != "" {
		variation = model.NewMix(game.MixedGame(*mixedGame))
	} else {
		variation = model.NewGame(game.LimitedGame(*limitedGame), game.FixedLimit)
	}

	play := play.NewPlay(variation, stake, table)

	players := []model.Player{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	stack := 1500.

	for i, player := range players {
		if i < size {
			play.Recv <- protocol.NewJoinTable(player, i, stack)
			play.Broadcast.Bind(player, &me)
		}
	}

	return play
}