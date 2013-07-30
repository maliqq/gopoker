package play

import (
	"fmt"
	"io"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/poker"
	"gopoker/poker/hand"
	"gopoker/protocol"
	"gopoker/protocol/message"
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

func (l *Logger) handle(msg *message.Message) {
	switch msg.Payload().(type) {
	case *message.RequireBet:
	case *message.RequireDiscard:

	case *message.DealCards:
		payload := msg.Envelope.DealCards

		if payload.GetType() == message.DealType_Board {
			l.log("Dealt %s [%s]\n",
				payload.GetType(),
				poker.BinaryCards(payload.Cards).PrintString(),
			)
		} else {
			l.log("Dealt %s [%s] to %d\n",
				payload.GetType(),
				poker.BinaryCards(payload.Cards).PrintString(),
				payload.GetPos(),
			)
		}

	case *message.MoveButton:
		payload := msg.Envelope.MoveButton
		l.log("Button is %d\n",
			payload.GetPos()+1,
		)

	case *message.AddBet:

		payload := msg.Envelope.AddBet
		betType := payload.Bet.GetType().String()

		l.log("Seat %d: %s\n",
			payload.GetPos(),
			model.NewBet(bet.Type(betType), payload.Bet.GetAmount()),
		)

	case *message.StreetStart:

		payload := msg.Envelope.StreetStart
		l.log("=== %s\n", payload.GetName())

	case *message.ShowHand:

		payload := msg.Envelope.ShowHand
		protoHand := payload.Hand
		hand := poker.Hand{
			Rank:   hand.Rank(protoHand.Rank.String()),
			High:   poker.BinaryCards(protoHand.High),
			Value:  poker.BinaryCards(protoHand.Value),
			Kicker: poker.BinaryCards(protoHand.Kicker),
		}
		l.log("Seat %d: shows [%s] (%s)\n", payload.GetPos(), poker.BinaryCards(payload.Cards).PrintString(), hand.PrintString())

	case *message.Winner:
		payload := msg.Envelope.Winner
		l.log("Seat %d: wins %.2f\n", payload.GetPos(), payload.GetAmount())

	default:
		l.log("got %s\n", msg.Envelope)
	}
}

func (l *Logger) log(format string, args ...interface{}) {
	fmt.Fprintf(l.Writer, format, args...)
}
