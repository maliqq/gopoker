package main

//
// start server node
//
import (
	"flag"
	"runtime"
)

import (
	"gopoker/model"
	"gopoker/server"
)

const (
	defaultConfigDir = "/etc/gopoker"
)

var (
	configDir = flag.String("config-dir", defaultConfigDir, "Config dir")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	model.LoadGames(*configDir)

	node := server.NewNode("localhost", *configDir)
	node.Start()
}
