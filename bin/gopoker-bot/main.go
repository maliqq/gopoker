package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

import (
	"gopoker/ai"
)

var (
	publisher = flag.String("publisher", "tcp://127.0.0.1:5555", "Node ZMQ publisher address")
	receiver  = flag.String("receiver", "tcp://127.0.0.1:5556", "Node ZMQ receiver address")
	roomID    = flag.String("roomid", "0", "Room ID")
	pos       = flag.Int("pos", 0, "Table position")
	stack     = flag.Float64("stack", 1000., "Stack amount")
)

var stdout = true

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	// logs
	if !stdout {
		w, err := os.Create(fmt.Sprintf("/var/log/gopoker/bot-%d", *pos))
		if err != nil {
			panic(err.Error())
		}
		defer w.Close()
		log.SetOutput(w)
	}

	id := fmt.Sprintf("player-%d", *pos)
	bot := ai.NewBot(id, *publisher, *receiver)
	bot.Join(*roomID, *pos, *stack)
	bot.Play()
}
