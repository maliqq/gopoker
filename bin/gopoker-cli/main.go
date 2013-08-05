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
	"gopoker/protocol/message"

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

	me := make(protocol.MessageChannel, 100)
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

func createPlay(me *protocol.MessageChannel) *play.Play {
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
			play.Broadcast.Bind(player, me)
			play.Recv <- message.NewJoinTable(string(player), i, stack)
		}
	}

	return play
}