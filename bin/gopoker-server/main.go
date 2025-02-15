package main

//
// start server node
//
import (
	"flag"
	"log"
	"runtime"
)

import (
	"gopoker/model"
	"gopoker/server"
)

const (
	defaultConfigDir = "/etc/gopoker"
	nodeConfigFile   = "node.json"
)

var (
	configDir = flag.String("config-dir", defaultConfigDir, "Config dir")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.SetPrefix("[server] ")

	flag.Parse()
	model.LoadGames(*configDir)

	var config *server.Config
	model.ReadConfig(*configDir, nodeConfigFile, &config)

	node := server.NewNode("localhost", config)
	node.Start()
}
