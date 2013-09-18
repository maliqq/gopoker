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
	"gopoker/event"
	"gopoker/event/message"

	"gopoker/model"
	"gopoker/model/game"

	"gopoker/client/cli"
	"gopoker/play"
)

var (
	logfile     = flag.String("logfile", "", "Log file path")
	betsize     = flag.Float64("betsize", 20., "Bet size")
	mixedGame   = flag.String("mix", "", "Mix to play")
	limitedGame = flag.String("game", "Texas", "Game to play")
)

const (
	defaultConfigDir = "/etc/gopoker"
)

var (
	configDir = flag.String("config-dir", defaultConfigDir, "Config dir")
)

func main() {
	flag.Parse()
	model.LoadGames(*configDir)

	if *logfile != "" {
		w, err := os.Create(*logfile)

		if err != nil {
			panic(err.Error())
		}

		defer w.Close()
		log.SetOutput(w)
	}

	me := make(event.Channel, 100)
	play := createPlay(&me)
	fmt.Printf("%s\n", play)

	play.Start()

	conn := &cli.Connection{
		Server: play,
	}

	for msg := range me {
		conn.Handle(msg)
	}
}

func createPlay(me *event.Channel) *play.Play {
	size := 3
	stake := model.NewStake(*betsize)
	//stake.WithAnte = true

	var variation model.Variation
	if *mixedGame != "" {
		variation = model.NewMix(game.MixedGame(*mixedGame), size)
	} else {
		variation = model.NewGame(game.LimitedGame(*limitedGame), game.FixedLimit, size)
	}

	play := play.NewPlay(variation, stake)

	players := []model.Player{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	stack := 1500.

	for i, player := range players {
		if i < size {
			play.Broadcast.Bind(player, me)
			play.Recv <- event.NewEvent(&message.JoinTable{player, i, stack})
		}
	}

	return play
}
