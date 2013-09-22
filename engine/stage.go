package engine

import (
  "log"
)

import (
  "gopoker/engine/stage"
)

type Stage struct {
  Type     stage.Type
  If       func() bool
  Before   func()
  Do       func()
  After    func()
  Notify   bool
}

func (process *Stage) Run() {
  if process.If != nil && !process.If() {
    return
  }

  log.Printf("[stage] start %s", process.Type)

  if process.Before != nil {
    process.Before()
  }

  if process.Do != nil {
    process.Do()
  }

  if process.After != nil {
    process.After()
  }
}

type Stages []Stage

func (processes Stages) Run() {
  for _, process := range processes {
    process.Run()
  }
}

