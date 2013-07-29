package play

import (
	"fmt"
	"io"
)

import (
	"gopoker/poker"
	"gopoker/poker/hand"
	"gopoker/protocol"
)

type Logger struct {
	Writer io.Writer
	Recv   protocol.MessageChannel
}

func NewLogger(writer io.Writer) *Logger {
	logger := &Logger{
		Writer: writer,
		Recv:   make(protocol.MessageChannel),
	}

	go logger.receive()

	return logger
}

func (l *Logger) receive() {
	for {
		msg := <-l.Recv
		l.handle(msg)
	}
}

func (l *Logger) handle(msg *protocol.Message) {
	switch msg.Payload().(type) {
	case protocol.RequireBet:
	case protocol.RequireDiscard:

	case protocol.DealCards:
		payload := msg.Envelope.DealCards

		if *payload.Type == protocol.DealType_Board {
			l.log("Dealt %s [%s]\n", payload.Type, poker.FromBinary(payload.Cards).PrintString())
		} else {
			l.log("Dealt %s [%s] to %d\n", payload.Type, poker.FromBinary(payload.Cards).PrintString(), payload.Pos)
		}

	case protocol.MoveButton:
		payload := msg.Envelope.MoveButton
		l.log("Button is %d\n", *payload.Pos+1)

	case protocol.AddBet:

		payload := msg.Envelope.AddBet
		l.log("Seat %d: %s\n", payload.Pos, payload.Bet)

	case protocol.StreetStart:

		payload := msg.Envelope.StreetStart
		l.log("=== %s\n", payload.Name)

	case protocol.ShowHand:

		payload := msg.Envelope.ShowHand
		handData := payload.Hand
		hand := poker.Hand{
			Rank:   hand.Rank(handData.Rank.String()),
			High:   poker.FromBinary(handData.High),
			Value:  poker.FromBinary(handData.Value),
			Kicker: poker.FromBinary(handData.Kicker),
		}
		l.log("Seat %d: shows [%s] (%s)\n", payload.Pos, poker.FromBinary(payload.Cards).PrintString(), hand.PrintString())

	case protocol.Winner:
		payload := msg.Envelope.Winner
		l.log("Seat %d: wins %.2f\n", payload.Pos, payload.Amount)

	default:
		l.log("got %#v\n", msg.Envelope)
	}
}

func (l *Logger) log(format string, args ...interface{}) {
	fmt.Fprintf(l.Writer, format, args...)
}
