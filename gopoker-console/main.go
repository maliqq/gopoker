package main

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

	"gopoker/play"
	"gopoker/play/command"

	_ "gopoker/client"
	"gopoker/client/console_client"
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

	go play.Run()
	play.Control <- command.NextDeal

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

	ids := []model.Id{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	stack := 1500.
	players := make([]*model.Player, 9)

	for i, id := range ids {
		player := model.NewPlayer(id)
		players[i] = player

		if i < size {
			play.Receive <- protocol.NewJoinTable(player, i, stack)
			play.Broadcast.Bind(player, me)
		}
	}

	return play
}
