package console_client

import (
	"bufio"
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

	"gopoker/play"
)

type Connection struct {
	Server *play.Play
}

func (c *Connection) Reply(msg *protocol.Message) {
	c.Server.Receive <- msg
}

func (c *Connection) Handle(msg *protocol.Message) {
	log.Println(console.Color(console.GREEN, fmt.Sprintf("[receive] %s", msg)))

	switch msg.Payload().(type) {
	case protocol.RequireBet:

		r := msg.Envelope.RequireBet

		fmt.Printf("%s\n", r)

		var newBet *model.Bet

		seat := c.Server.Table.Seat(r.Pos)

		for newBet == nil {
			newBet = readBet(r.Call, r.Call-seat.Bet)

			err := newBet.Validate(seat, r.BetRange)
			if err != nil {
				fmt.Println(err.Error())
				newBet = nil
			}
		}

		if newBet != nil {
			c.Reply(protocol.NewAddBet(r.Pos, newBet))
		}

	case protocol.RequireDiscard:

		r := msg.Envelope.RequireDiscard

		seat := c.Server.Table.Seat(r.Pos)

		fmt.Printf("Your cards: [%s]\n", c.Server.Deal.Pocket(seat.Player))

		var cards *poker.Cards
		for cards == nil {
			cards = readCards()
		}

		c.Reply(protocol.NewDiscardCards(r.Pos, cards))

	case protocol.DealCards:

		payload := msg.Envelope.DealCards

		if payload.Type.IsBoard() {
			fmt.Printf("Dealt %s %s\n", payload.Type, payload.Cards.ConsoleString())
		} else {
			fmt.Printf("Dealt %s %s to %d\n", payload.Type, payload.Cards.ConsoleString(), payload.Pos)
		}

	case protocol.MoveButton:

		payload := msg.Envelope.MoveButton

		fmt.Printf("Button is %d\n", payload.Pos+1)

	case protocol.AddBet:

		payload := msg.Envelope.AddBet

		player := c.Server.Table.Player(payload.Pos)

		fmt.Printf("Player %s: %s\n", player, payload.Bet)

	case protocol.PotSummary:

		payload := msg.Envelope.PotSummary

		fmt.Printf("Pot size: %.2f\nBoard: %s\n", payload.Amount, c.Server.Deal.Board.ConsoleString())

	case protocol.ShowHand:

		payload := msg.Envelope.ShowHand

		player := c.Server.Table.Player(payload.Pos)

		fmt.Printf("Player %s has %s (%s)\n", player, payload.Cards.ConsoleString(), payload.Hand.HumanString())

	case protocol.Winner:

		payload := msg.Envelope.Winner

		fmt.Printf("Player %s won %.2f\n", payload.Player, payload.Amount)

	case protocol.ChangeGame:

		payload := msg.Envelope.ChangeGame

		fmt.Printf("Game changed to %s %s\n", payload.Type, payload.Limit)
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
