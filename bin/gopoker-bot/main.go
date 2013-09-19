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
	group     = flag.Int("group", 0, "Bots group")
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

	if *group > 0 {
		log.Printf("starting %d bots...", *group)
		for i := 0; i < *group; i++ {
			go startBot(i)
		}
		for {
		}
	} else {
		startBot(*pos)
	}
}

func startBot(i int) {
	id := fmt.Sprintf("player-%d", i)
	log.Printf("spawning bot %s", id)
	bot := ai.NewBot(id, *roomID, *publisher, *receiver)
	bot.Join(i, *stack)
	bot.Play()
}
