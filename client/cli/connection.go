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
	"gopoker/poker"
	"gopoker/poker/hand"
	"gopoker/protocol/message"
	"gopoker/util/console"

	"gopoker/model"
	"gopoker/model/bet"

	"gopoker/play"
)

// Connection - console connection
type Connection struct {
	Server *play.Play
}

// Reply - reply to play
func (c *Connection) Reply(msg *message.Message) {
	c.Server.Recv <- msg
}

// Handle - handle protocol message
func (c *Connection) Handle(msg *message.Message) {
	log.Println(console.Color(console.Green, fmt.Sprintf("[receive] %s", msg)))

	switch msg.Payload().(type) {
	case *message.RequireBet:

		r := msg.Envelope.RequireBet

		fmt.Printf("%s\n", r)

		var newBet *model.Bet

		pos := int(r.GetPos())
		seat := c.Server.Table.Seat(pos)

		for newBet == nil {
			betRange := &bet.Range{
				Call: r.BetRange.GetCall(),
				Min:  r.BetRange.GetMin(),
				Max:  r.BetRange.GetMax(),
			}
			newBet = readBet(betRange.Call, betRange.Call-seat.Bet)

			err := newBet.Validate(seat, betRange)
			if err != nil {
				fmt.Println(err.Error())
				newBet = nil
			}
		}

		if newBet != nil {
			pos := int(r.GetPos())
			c.Reply(message.NotifyAddBet(pos, newBet.Proto()))
		}

	case *message.RequireDiscard:

		r := msg.Envelope.RequireDiscard
		pos := int(r.GetPos())

		seat := c.Server.Table.Seat(pos)

		fmt.Printf("Your cards: [%s]\n", c.Server.Deal.Pocket(seat.Player))

		var cards poker.Cards
		for cards == nil {
			cards = readCards()
		}

		c.Reply(message.NotifyDiscardCards(pos, cards.Proto()))

	case *message.DealCards:

		payload := msg.Envelope.DealCards

		if payload.GetType() == message.DealType_Board {
			fmt.Printf("Dealt %s %s\n", payload.Type, poker.BinaryCards(payload.Cards).ConsoleString())
		} else {
			fmt.Printf("Dealt %s %s to %d\n", payload.Type, poker.BinaryCards(payload.Cards).ConsoleString(), payload.Pos)
		}

	case *message.MoveButton:

		payload := msg.Envelope.MoveButton
		pos := int(payload.GetPos())

		fmt.Printf("Button is %d\n", pos+1)

	case *message.AddBet:

		payload := msg.Envelope.AddBet
		pos := int(payload.GetPos())
		player := c.Server.Table.Player(pos)
		betType := payload.Bet.GetType().String()

		fmt.Printf("Player %s: %s\n", player, model.NewBet(bet.Type(betType), payload.Bet.GetAmount()))

	case *message.BettingComplete:

		payload := msg.Envelope.BettingComplete
		pot := int(payload.GetPot())

		fmt.Printf("Pot size: %.2f\nBoard: %s\n", pot, c.Server.Deal.Board.ConsoleString())

	case *message.ShowHand:

		payload := msg.Envelope.ShowHand
		pos := int(payload.GetPos())
		player := c.Server.Table.Player(pos)

		handData := payload.Hand
		hand := poker.Hand{
			Rank:   hand.Rank(handData.Rank.String()),
			High:   poker.BinaryCards(handData.High),
			Value:  poker.BinaryCards(handData.Value),
			Kicker: poker.BinaryCards(handData.Kicker),
		}

		fmt.Printf("Player %s has %s (%s)\n", player, poker.BinaryCards(payload.Cards).ConsoleString(), hand.PrintString())

	case *message.Winner:

		payload := msg.Envelope.Winner
		pos := int(payload.GetPos())
		player := c.Server.Table.Player(pos)
		amount := payload.GetAmount()

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
