package main

//
// start server node
//
import (
	"flag"
	"runtime"
)

import (
	"gopoker/server"
)

var (
	nodeConfigFile = flag.String("config-file", "", "Node config file")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	node := server.NewNode("localhost", *nodeConfigFile)
	node.Start()
}
