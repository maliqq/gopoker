package process

import "log"

import (
  "gopoker/engine/street"
)

type Street struct {
  street.Type
  Stages
}

func (process *Street) Start() {
  log.Printf("[street] start %s", process.Type)
  
  process.Stages.Start()
}

func (process *Street) Stop() {
  log.Printf("[street] stop %s", process.Type)
}
