package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

import (
	"gopoker/event"

	"gopoker/message"
	"gopoker/poker"

	"gopoker/model"
	"gopoker/model/deal"

	"gopoker/engine"
)

// Connection - console connection
type Session struct {
	Instance *engine.Instance
}

// Reply - reply to play
func (s *Session) Send(msg message.Message) {
	s.Instance.Recv <- event.New(msg)
}

func (s *Session) Start(recv event.Channel) {
	for n := range recv {
		switch msg := n.Message.(type) {
		case *message.RequireBet:

			fmt.Printf("%s\n", msg)

			var newBet *model.Bet

			pos := msg.Pos
			seat := s.Instance.Table.Seat(pos)

			for newBet == nil {
				betRange := msg.Range
				newBet = readBet(betRange.Call, betRange.Call-seat.Bet)

				err := newBet.Validate(seat, betRange)
				if err != nil {
					fmt.Println(err.Error())
					newBet = nil
				}
			}

			if newBet != nil {
				s.Send(&message.AddBet{pos, newBet})
			}

		case *message.RequireDiscard:

			pos := msg.Pos
			seat := s.Instance.Table.Seat(pos)

			d := s.Instance.Deal()
			fmt.Printf("Your cards: [%s]\n", d.Pocket(seat.Player))

			var cards poker.Cards
			for cards == nil {
				cards = readCards()
			}

			s.Send(&message.DiscardCards{pos, cards})

		case *message.DealCards:

			if msg.Type == deal.Board {
				fmt.Printf("Dealt %s %s\n", msg.Type, msg.Cards.ConsoleString())
			} else {
				fmt.Printf("Dealt %s %s to %d\n", msg.Type, msg.Cards.ConsoleString(), msg.Pos)
			}

		case *message.MoveButton:

			pos := msg.Pos

			fmt.Printf("Button is %d\n", pos+1)

		case *message.AddBet:

			pos := msg.Pos
			player := s.Instance.Table.Player(pos)

			fmt.Printf("Player %s: %s\n", player, msg.Bet)

		case *message.BettingComplete:

			pot := msg.Pot

			d := s.Instance.Deal()
			fmt.Printf("Pot size: %.2f\nBoard: %s\n", pot, d.Board.ConsoleString())

		case *message.ShowHand:

			pos := msg.Pos
			hand := msg.Hand
			player := s.Instance.Table.Player(pos)

			fmt.Printf("Player %s has %s (%s)\n", player, msg.Cards.ConsoleString(), hand.PrintString())

		case *message.Winner:

			pos := msg.Pos
			amount := msg.Amount
			player := s.Instance.Table.Player(pos)

			fmt.Printf("Player %s won %.2f\n", player, amount)

		}
	}
}

func readLine() string {
	fmt.Print(">>> ")
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimRight(str, "\n")
	return str
}

func readBet(call float64, toCall float64) *model.Bet {
	var b *model.Bet

	for b == nil {
		b = parseBet(call, toCall, readLine())
	}

	return b
}

func readCards() poker.Cards {
	var cards poker.Cards
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

func parseBet(call float64, toCall float64, str string) *model.Bet {
	var b *model.Bet

	switch str {
	case "":
		if toCall == 0. { // nothing to call
			b = model.NewCheck()
		} else {
			b = model.NewCall(call)
		}

	case "fold":
		b = model.NewFold()

	case "check":
		b = model.NewCheck()

	case "call":
		b = model.NewCall(call)

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
