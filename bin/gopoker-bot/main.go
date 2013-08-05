package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

import (
	"gopoker/ai"
)

var (
	rpcAddr = flag.String("rpc", ":8081", "Node RPC address")
	zmqAddr = flag.String("zmq", "tcp://127.0.0.1:5555", "Node ZMQ address")
	roomID  = flag.String("roomid", "0", "Room ID")
	pos     = flag.Int("pos", 0, "Table position")
	stack   = flag.Float64("stack", 1000., "Stack amount")
)

var stdout = true

func main() {
	flag.Parse()

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
	bot := ai.NewBot(id, *rpcAddr, *zmqAddr)
	bot.Join(*roomID, *pos, *stack)
	bot.Play()
}