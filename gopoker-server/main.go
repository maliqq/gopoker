package main

import (
  "flag"
  "gopoker/server"
  "runtime"
)

var (
  apiAddr = flag.String("api", ":8080", "HTTP address")
  rpcAddr = flag.String("rpc", ":8081", "RPC address")
)

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  flag.Parse()

  node := server.CreateNode("localhost", *apiAddr, *rpcAddr)
  node.Start()
}
