package model

import (
	"encoding/json"
	"os"
	"path"
)

import (
	"github.com/golang/glog"
)

// ReadConfig - read config file
func ReadConfig(configDir, filename string, result interface{}) {
	filepath := path.Join(configDir, filename)
	f, err := os.Open(filepath)

	if err != nil {
		glog.Fatalf("Can't open %s: %s", filepath, err)
	}

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&result)
	if err != nil {
		glog.Fatalf("Can't decode %s: %s", filepath, err)
	}
}
