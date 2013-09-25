package server

import (
	"io"
	"log"
)

import (
	"gopoker/event"
	"gopoker/event/message"
	"gopoker/model/deal"
)

// Logger - play file logger
type Logger struct {
	*log.Logger
}

// NewLogger - create new logger
func NewLogger(writer io.Writer) *Logger {
	return &Logger{
		Logger: log.New(writer, "", log.LstdFlags),
	}
}

func (l *Logger) HandleEvent(event *exchange.Event) {
	switch msg := event.Message.(type) {
	case *message.RequireBet:
	case *message.RequireDiscard:

	case *message.DealCards:

		if msg.Type == deal.Board {
			l.Printf("Dealt %s [%s]\n",
				msg.Type,
				msg.Cards.PrintString(),
			)
		} else {
			l.Printf("Dealt %s [%s] to %d\n",
				msg.Type,
				msg.Cards.PrintString(),
				msg.Pos,
			)
		}

	case *message.MoveButton:

		l.Printf("Button is %d\n",
			msg.Pos+1,
		)

	case *message.AddBet:

		l.Printf("Seat %d: %s\n",
			msg.Pos,
			msg.Bet,
		)

	case *message.StreetStart:

		l.Printf("=== %s\n", msg.Name)

	case *message.ShowHand:

		l.Printf("Seat %d: shows [%s] (%s)\n",
			msg.Pos,
			msg.Cards.PrintString(),
			msg.Hand.PrintString(),
		)

	case *message.Winner:

		l.Printf("Seat %d: wins %.2f\n",
			msg.Pos,
			msg.Amount,
		)

	default:

		l.Printf("got %s\n", event)
	}
}
