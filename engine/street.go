package engine

import (
	"log"
)

import (
  "gopoker/engine/street"
)

type Street struct {
  street.Type
  Stages
}

func (process *Street) Start() {
  log.Printf("[street] %s", process.Type)

  process.Stages.Run()
}

func (process *Street) Stop() {
}
