package main

//
// start server node
//
import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"runtime"
)

import (
	"gopoker/server"
)

var (
	config = flag.String("config", "/etc/gopoker.json", "Config file path")
)

func readConfig() *server.Config {
	f, err := os.Open(*config)
	if err != nil {
		log.Fatal("read config error", err)
	}
	var config server.Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("parse config error", err)
	}

	return &config
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	node := server.NewNode("localhost", readConfig())
	node.Start()
}
