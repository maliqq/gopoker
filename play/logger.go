package play

import (
	"log"
)

import (
	"gopoker/protocol"
)

type Logger struct {
	Recv protocol.MessageChannel
}

func NewLogger() *Logger {
	logger := &Logger{
		Recv: make(protocol.MessageChannel),
	}

	go logger.receive()

	return logger
}

func (l *Logger) receive() {
	for {
		msg := <-l.Recv
		log.Printf("%s", msg)
	}
}
