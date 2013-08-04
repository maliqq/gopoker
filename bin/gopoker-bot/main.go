package main

import (
  "flag"
)

import (
  "gopoker/ai"
)

var (
  rpcAddr = flag.String("rpc", ":8081", "Node RPC address")
  zmqAddr = flag.String("zmq", "tcp://localhost:5555", "Node ZMQ address")
  roomID = flag.String("roomid", "0", "Room ID")
  pos = flag.Int("pos", 0, "Table position")
  amount = flag.Float64("amount", 1000., "Stack amount")
)

func main() {
  flag.Parse()
  
  bot := ai.NewBot(*rpcAddr, *zmqAddr)
  bot.Join(*roomID, *pos, *amount)
  bot.Play()
}
