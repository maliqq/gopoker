package process

import "log"

type Stage struct {
  Name string
  If     func() bool
  Before func()
  Run    func()
  After  func()
  Notify bool
}

func (process *Stage) Stop() {
  log.Printf("[stage] stop %s", process.Name)
}

func (process *Stage) Start() {
  if process.If != nil && !process.If() {
    return
  }

  log.Printf("[stage] start %s", process.Name)

  if process.Before != nil {
    process.Before()
  }

  if process.Run != nil {
    process.Run()
  }

  if process.After != nil {
    process.After()
  }
}

type Stages []Stage

func (processes Stages) Start() {
  for _, process := range processes {
    process.Start()
  }
}
