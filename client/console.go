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
	"gopoker/poker"
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/deal"
	"gopoker/model/game"
	runner "gopoker/play"
	"gopoker/play/command"
	"gopoker/play/context"
	"gopoker/protocol"
	"gopoker/util/console"
)

var (
	logfile    = flag.String("logfile", "", "Log file path")
	betsize    = flag.Float64("betsize", 20., "Bet size")
	gametoplay = flag.String("game", "texas", "Game to play")
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

	go runner.Run(play)
	play.Control <- command.NextDeal

	for {
		select {
		case msg := <-me:

			log.Printf("%sreceived %s%s\n", console.GREEN, msg.String(), console.RESET)

			switch msg.Payload.(type) {
			case protocol.RequireBet:

				r := msg.Payload.(protocol.RequireBet)

				fmt.Printf("%s\n", r)

				var newBet *bet.Bet
				for newBet == nil {
					newBet, _ = readBet(&r)

					err := play.Betting.ValidateBet(play.Table.Seat(r.Pos), newBet)
					if err != nil {
						fmt.Println(err.Error())
						newBet = nil
					}
				}

				play.Receive <- protocol.NewAddBet(r.Pos, newBet)

			case protocol.RequireDiscard:

				r := msg.Payload.(protocol.RequireDiscard)

				seat := play.Table.Seat(r.Pos)

				fmt.Printf("Your cards: [%s]\n", play.Deal.Pocket(seat.Player))

				var cards *poker.Cards
				for cards == nil {
					cards = readCards()
				}

				play.Receive <- protocol.NewDiscardCards(r.Pos, cards)

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

				player := play.Table.Player(payload.Pos)

				fmt.Printf("Player %s: %s\n", player, payload.Bet)

			case protocol.PotSummary:

				payload := msg.Payload.(protocol.PotSummary)

				fmt.Printf("Pot size: %.2f\n", payload.Amount)

			case protocol.ShowHand:

				payload := msg.Payload.(protocol.ShowHand)

				player := play.Table.Player(payload.Pos)

				fmt.Printf("Player %s has %s (%s)\n", player, payload.Cards.ConsoleString(), payload.Hand.ConsoleString())

			case protocol.Winner:

				payload := msg.Payload.(protocol.Winner)

				fmt.Printf("Player %s won %.2f\n", payload.Player, payload.Amount)
			}
		}
	}
	fmt.Println("bye")
}

func createPlay(me protocol.MessageChannel) *context.Play {
	size := 3
	stake := game.NewStake(*betsize)
	//stake.WithAnte = true
	g := model.NewGame(game.LimitedGame(*gametoplay), game.FixedLimit, stake)
	table := model.NewTable(size)
	play := context.NewPlay(g, table)
	fmt.Printf("play.Game()=%#v", play.Game())

	ids := []model.Id{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	stack := 1500.
	players := make([]*model.Player, 9)

	for i, id := range ids {
		player := model.NewPlayer(id)
		players[i] = player

		if i < size {
			table.AddPlayer(player, i, stack)
			play.Broadcast.Bind(player, me)
		}
	}

	return play
}

func readLine() string {
	fmt.Print(">>> ")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimRight(str, "\n")
	return str
}

func readBet(r *protocol.RequireBet) (*bet.Bet, string) {
	var b *bet.Bet
	var str string

	for b == nil {
		str = readLine()

		if str == "exit" {
			return nil, "exit"
		}

		b = parseBet(r, str)
	}

	return b, str
}

func readCards() *poker.Cards {
	var cards *poker.Cards
	for cards == nil {
		str := readLine()
		var err error
		cards, err = poker.ParseCards(str)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}
	return cards
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
				fmt.Printf("error: %s\n", err.Error())
			}
		}
	}

	return b
}
