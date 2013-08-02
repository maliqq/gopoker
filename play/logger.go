package play

import (
	"io"
	"log"
)

import (
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/poker"
	"gopoker/poker/hand"
	"gopoker/protocol"
	"gopoker/protocol/message"
)

// Logger - play file logger
type Logger struct {
	*log.Logger
	Recv protocol.MessageChannel
}

// NewLogger - create new logger
func NewLogger(writer io.Writer) *Logger {
	logger := &Logger{
		Logger: log.New(writer, "", log.LstdFlags),
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
			l.Printf("Dealt %s [%s]\n",
				payload.GetType(),
				poker.BinaryCards(payload.Cards).PrintString(),
			)
		} else {
			l.Printf("Dealt %s [%s] to %d\n",
				payload.GetType(),
				poker.BinaryCards(payload.Cards).PrintString(),
				payload.GetPos(),
			)
		}

	case *message.MoveButton:
		payload := msg.Envelope.MoveButton

		l.Printf("Button is %d\n",
			payload.GetPos()+1,
		)

	case *message.AddBet:
		payload := msg.Envelope.AddBet
		betType := payload.Bet.GetType().String()

		l.Printf("Seat %d: %s\n",
			payload.GetPos(),
			model.NewBet(bet.Type(betType), payload.Bet.GetAmount()),
		)

	case *message.StreetStart:
		payload := msg.Envelope.StreetStart

		l.Printf("=== %s\n", payload.GetName())

	case *message.ShowHand:
		payload := msg.Envelope.ShowHand
		protoHand := payload.Hand
		hand := poker.Hand{
			Rank:   hand.Rank(protoHand.Rank.String()),
			High:   poker.BinaryCards(protoHand.High),
			Value:  poker.BinaryCards(protoHand.Value),
			Kicker: poker.BinaryCards(protoHand.Kicker),
		}

		l.Printf("Seat %d: shows [%s] (%s)\n",
			payload.GetPos(),
			poker.BinaryCards(payload.Cards).PrintString(),
			hand.PrintString(),
		)

	case *message.Winner:
		payload := msg.Envelope.Winner

		l.Printf("Seat %d: wins %.2f\n",
			payload.GetPos(),
			payload.GetAmount(),
		)

	default:
		l.Printf("got %s\n", msg.Envelope)
	}
}
