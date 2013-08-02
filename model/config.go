package model

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

// ReadConfig - read config file
func ReadConfig(configDir, filename string, result interface{}) {
	filepath := path.Join(configDir, filename)
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
