package engine

import (
	"log"
)

import (
	"gopoker/engine/stage"
)

type Stage struct {
	Type   stage.Type
	If     func() bool
	Before func()
	After  func()
	Notify bool
}

func (process *Stage) do(doFunc func()) {
	if process.If != nil && !process.If() {
		return
	}

	log.Printf("[stage] start %s", process.Type)

	if process.Before != nil {
		process.Before()
	}

	if doFunc != nil {
		doFunc()
	}

	if process.After != nil {
		process.After()
	}
}

func (process *Stage) String() string {
	return string(process.Type)
}
