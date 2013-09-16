package play

import (
	"io"
	"log"
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
