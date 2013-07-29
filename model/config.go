package model

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path"
)

const (
	DefaultConfigDir = "/etc/gopoker"
	GamesConfigFile  = "games.json"
	MixesConfigFile  = "mixes.json"
)

var ConfigDir = flag.String("config-dir", DefaultConfigDir, "Config dir")

func ReadConfig(filename string, result interface{}) {
	filepath := path.Join(*ConfigDir, filename)
	f, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Can't open %s: %s", filepath, err)
	}

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&result)
	if err != nil {
		log.Fatalf("Can't decode %s: %s", filepath, err)
	}
}

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
}
