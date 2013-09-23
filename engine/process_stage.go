package engine

type StageProcess interface {
  Run() (bool, bool)
}

type StageDo struct {
  Stage
  Do func()
}

type StageExit struct {
  Stage
  Do func(chan bool)
}

type StageSkip struct {
  Stage
  Do func(chan bool)
}

func (stage StageDo) Run() (bool, bool) {
  stage.do(stage.Do)
  return false, false
}

func (stage StageExit) Run() (bool, bool) {
  exit := make(chan bool)
  
  var doExit bool
  go func() {
    _, doExit = <-exit
  }()
  stage.do(func() {
    stage.Do(exit)
  })
  
  return doExit, false
}

func (stage StageSkip) Run() (bool, bool) {
  skip := make(chan bool)

  var doSkip bool
  go func() {
    _, doSkip = <-skip
  }()
  stage.do(func() {
    stage.Do(skip)
  })

  return false, doSkip
}

type Stages []StageProcess
func (stages Stages) Run() bool {
  for _, stage := range stages {
    doExit, doSkip := stage.Run()
    if doExit {
      return false
    }
    if doSkip {
      break
    }
  }

  return true
}
