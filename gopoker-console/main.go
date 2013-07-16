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
	"gopoker/protocol"
	"gopoker/util/console"

	"gopoker/model"
	"gopoker/model/game"
	
	"gopoker/play"
	"gopoker/play/command"
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
	go play.Run()
	play.Control <- command.NextDeal

	fmt.Printf("%s\n", play)

Loop:
	for {
		select {
		case msg := <-me:

			log.Println(console.Color(console.GREEN, fmt.Sprintf("[receive] %s", msg)))

			switch msg.Payload.(type) {
			case protocol.RequireBet:

				r := msg.Payload.(protocol.RequireBet)

				fmt.Printf("%s\n", r)

				var newBet *model.Bet
			ReadBetLoop:
				for newBet == nil {
					var cmd string
					newBet, cmd = readBet(&r)
					switch cmd {
					case "exit":
						play.Control <- command.Exit
						break Loop

					case "sitout":
						play.Receive <- protocol.NewMessage(protocol.SitOut{Pos: r.Pos})
						break ReadBetLoop
					}

					err := newBet.Validate(play.Table.Seat(r.Pos), play.Betting.Required.RequireBet)
					if err != nil {
						fmt.Println(err.Error())
						newBet = nil
					}
				}

				if newBet != nil {
					play.Receive <- protocol.NewAddBet(r.Pos, newBet)
				}

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

				if payload.Type.IsBoard() {
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

				fmt.Printf("Player %s has %s (%s)\n", player, payload.Cards.ConsoleString(), payload.Hand.HumanString())

			case protocol.Winner:

				payload := msg.Payload.(protocol.Winner)

				fmt.Printf("Player %s won %.2f\n", payload.Player, payload.Amount)

			case protocol.ChangeGame:

				payload := msg.Payload.(protocol.ChangeGame)

				fmt.Printf("Game changed to %s %s\n", payload.Type, payload.Limit)
			}
		}
	}
	fmt.Println("bye")
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

func readLine() string {
	fmt.Print(">>> ")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimRight(str, "\n")
	return str
}

func readBet(r *protocol.RequireBet) (*model.Bet, string) {
	var b *model.Bet
	var str string

	for b == nil {
		str = readLine()

		if str == "exit" || str == "sitout" {
			return nil, str
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

func parseBet(r *protocol.RequireBet, str string) *model.Bet {
	var b *model.Bet

	switch str {
	case "":
		if r.Call == 0. {
			b = model.NewCheck()
		} else {
			b = model.NewCall(r.Call)
		}

	case "fold":
		b = model.NewFold()

	case "check":
		b = model.NewCheck()

	case "call":
		b = model.NewCall(r.Call)

	default:
		parts := strings.Split(str, " ")

		var amountString string
		if len(parts) == 1 {
			amountString = parts[0]
		} else if len(parts) == 2 && parts[0] == "raise" {
			amountString = parts[1]
		}

		if amountString != "" {
			amount, err := strconv.ParseFloat(amountString, 64)

			if err == nil {
				b = model.NewRaise(amount)
			} else {
				fmt.Printf("error: %s\n", err.Error())
			}
		}
	}

	return b
}
