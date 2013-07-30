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
	DefaultConfigDir = "/etc/gopoker"
)

var (
	ConfigDir = flag.String("config-dir", DefaultConfigDir, "Config dir")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	model.LoadGames(*ConfigDir)

	node := server.NewNode("localhost", *ConfigDir)
	node.Start()
}
