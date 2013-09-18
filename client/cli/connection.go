package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

import (
	"gopoker/event"
	"gopoker/event/message"
	"gopoker/poker"
	"gopoker/util"

	"gopoker/model"
	"gopoker/model/deal"

	"gopoker/play"
)

// Connection - console connection
type Connection struct {
	Server *play.Play
}

// Reply - reply to play
func (c *Connection) Reply(msg message.Message) {
	c.Server.Recv <- event.NewEvent(msg)
}

// Handle - handle protocol message
func (c *Connection) Handle(event *event.Event) {
	log.Println(util.Color(util.Green, fmt.Sprintf("[receive] %s", event)))

	switch msg := event.Message.(type) {
	case *message.RequireBet:

		fmt.Printf("%s\n", msg)

		var newBet *model.Bet

		pos := msg.Pos
		seat := c.Server.Table.Seat(pos)

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
			c.Reply(&message.AddBet{pos, newBet})
		}

	case *message.RequireDiscard:

		pos := msg.Pos
		seat := c.Server.Table.Seat(pos)

		fmt.Printf("Your cards: [%s]\n", c.Server.Deal.Pocket(seat.Player))

		var cards poker.Cards
		for cards == nil {
			cards = readCards()
		}

		c.Reply(&message.DiscardCards{pos, cards})

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
		player := c.Server.Table.Player(pos)

		fmt.Printf("Player %s: %s\n", player, msg.Bet)

	case *message.BettingComplete:

		pot := msg.Pot

		fmt.Printf("Pot size: %.2f\nBoard: %s\n", pot, c.Server.Deal.Board.ConsoleString())

	case *message.ShowHand:

		pos := msg.Pos
		hand := msg.Hand
		player := c.Server.Table.Player(pos)

		fmt.Printf("Player %s has %s (%s)\n", player, msg.Cards.ConsoleString(), hand.PrintString())

	case *message.Winner:

		pos := msg.Pos
		amount := msg.Amount
		player := c.Server.Table.Player(pos)

		fmt.Printf("Player %s won %.2f\n", player, amount)
		/*
			case *message.ChangeGame:

				payload := msg.Envelope.ChangeGame

				fmt.Printf("Game changed to %s %s\n", payload.Type, payload.Limit)
		*/
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
