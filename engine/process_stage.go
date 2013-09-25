package engine

import (
	"github.com/golang/glog"
)

type StageProcess interface {
	Run() (bool, bool)
	String() string
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

type StageStrategy []StageProcess

func (stages StageStrategy) Run() bool {
	for _, stage := range stages {
		doExit, doSkip := stage.Run()
		if doExit {
			glog.Infof("[stage] exit %s", stage.String())
			return false
		}
		if doSkip {
			glog.Infof("[stage] skip %s", stage.String())
			break
		}
	}

	return true
}
