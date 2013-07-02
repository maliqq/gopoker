package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/deal"
	"gopoker/model/game"
	"gopoker/play"
	"gopoker/play/context"
	_ "gopoker/play/stage"
	"gopoker/protocol"
	"gopoker/util/console"
)

var (
	logfile = flag.String("logfile", "", "Log file path")
	betsize = flag.Float64("betsize", 20., "Bet size")
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
	playContext := createPlayContext(me)

	fmt.Printf("%s\n", playContext)

	go play.Start(playContext)

	for {
		select {
		case msg := <-me:

			log.Printf("%sreceived %s%s\n", console.GREEN, msg.String(), console.RESET)

			switch msg.Payload.(type) {
			case protocol.RequireBet:

				r := msg.Payload.(protocol.RequireBet)

				fmt.Printf("Require %s\n", r)

				var newBet *bet.Bet
				for newBet == nil {
					newBet = readBet(&r)

					err := context.ValidateBet(&r, playContext.Table.Seat(r.Pos), newBet)
					if err != nil {
						fmt.Println(err.Error())
						newBet = nil
					}
				}

				playContext.Receive <- protocol.NewAddBet(r.Pos, newBet)

			case protocol.DealCards:

				payload := msg.Payload.(protocol.DealCards)

				if payload.Type == deal.Board {
					fmt.Printf("Dealt %s %s\n", payload.Type, payload.Cards.ConsoleString())
				} else {
					fmt.Printf("Dealt %s %s to %d\n", payload.Type, payload.Cards.ConsoleString(), payload.Pos)
				}

			case protocol.MoveButton:

				payload := msg.Payload.(protocol.MoveButton)

				fmt.Printf("Button is %d\n", payload.Pos+1)

			case protocol.AddBet:

				payload := msg.Payload.(protocol.AddBet)

				player := playContext.Table.Player(payload.Pos)

				fmt.Printf("Player %s: %s\n", player, payload.Bet)

			case protocol.ShowHand:

				payload := msg.Payload.(protocol.ShowHand)

				player := playContext.Table.Player(payload.Pos)

				fmt.Printf("Player %s has %s (%s)\n", player, payload.Cards.ConsoleString(), payload.Hand.ConsoleString())
			}
		}
	}
}

func createPlayContext(me protocol.MessageChannel) *context.Play {
	size := 3
	game := model.NewGame(game.Texas, game.FixedLimit, game.NewStake(*betsize))
	table := model.NewTable(size)
	playContext := context.NewPlay(game, table)

	ids := []model.Id{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	stack := 1500.
	players := make([]*model.Player, 9)

	for i, id := range ids {
		player := model.NewPlayer(id)
		players[i] = player

		if i < size {
			table.AddPlayer(player, i, stack)
			playContext.Broadcast.Bind(player, me)
		}
	}

	return playContext
}

func readBet(r *protocol.RequireBet) *bet.Bet {
	var b *bet.Bet

	for b == nil {
		fmt.Print(">>> ")
		betString, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		b = parseBet(r, strings.TrimRight(betString, "\n"))
	}

	return b
}

func parseBet(r *protocol.RequireBet, betString string) *bet.Bet {
	var b *bet.Bet

	switch betString {
	case "":
		if r.Call == 0. {
			b = bet.NewCheck()
		} else {
			b = bet.NewCall(r.Call)
		}

	case "fold":
		b = bet.NewFold()

	case "check":
		b = bet.NewCheck()

	case "call":
		b = bet.NewCall(r.Call)

	default:
		parts := strings.Split(betString, " ")

		var amountString string
		if len(parts) == 1 {
			amountString = parts[0]
		} else if len(parts) == 2 && parts[0] == "raise" {
			amountString = parts[1]
		}

		if amountString != "" {
			amount, err := strconv.ParseFloat(amountString, 64)

			if err == nil {
				b = bet.NewRaise(amount)
			} else {
				fmt.Printf("error: %s", err.Error())
			}
		}
	}

	return b
}
