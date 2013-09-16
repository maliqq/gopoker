package play

import (
	"io"
	"log"
)

import (
	"gopoker/exch"
	"gopoker/exch/message"
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/poker"
	"gopoker/poker/hand"
)

// Logger - play file logger
type Logger struct {
	*log.Logger
	Recv exch.MessageChannel
}

// NewLogger - create new logger
func NewLogger(writer io.Writer) *Logger {
	logger := &Logger{
		Logger: log.New(writer, "", log.LstdFlags),
		Recv:   make(exch.MessageChannel),
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
