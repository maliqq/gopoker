package engine

import (
	"fmt"
	"time"
)

import (
	"gopoker/engine/mode"
	"gopoker/engine/street"
)

type State string

const (
	Waiting State = "waiting"
	Active  State = "active"
	Paused  State = "paused"
	Closed  State = "closed"
)

type InstanceStateChange struct {
	Timeout  int
	NewState State
}

type Instance struct {
	*Gameplay
	State       State
	StateChange chan InstanceStateChange
	Street      street.Type
	Mode        mode.Type
}

func NewInstance(context *Context) *Instance {
	instance := Instance{
		State:       Waiting,
		StateChange: make(chan InstanceStateChange),
		Gameplay:    NewGameplay(context),
	}

	return &instance
}

func (instance *Instance) doStart() {
	fmt.Println("start...")
	instance.State = Active
}

func (instance *Instance) doPause() {
	fmt.Println("pause...")
	instance.State = Paused
}

func (instance *Instance) doResume() {
	fmt.Println("resume...")
	instance.State = Waiting
}

func (instance *Instance) doStop() {
	fmt.Println("stop...")
	instance.State = Closed
}

func (instance *Instance) processStateChange(event InstanceStateChange) bool {
	// check for correct transition
	switch instance.State {
	case Paused, Closed:
		if event.NewState != Waiting {
			return false
		}
	case Waiting:
		if event.NewState != Active || event.NewState != Closed {
			return false
		}
	}

	if timeout := event.Timeout; timeout != 0 {
		<-time.After(time.Duration(timeout) * time.Second)
	}
	instance.State = event.NewState

	return true
}

func (instance *Instance) Run() {
RunLoop:
	for {
		select {
		case event := <-instance.StateChange:
			if !instance.processStateChange(event) {
				break
			}

			switch event.NewState {
			case Active:
				instance.doStart()
			case Paused:
			PauseLoop:
				for {
					select {
					case event := <-instance.StateChange:
						if instance.processStateChange(event) {
							break PauseLoop
						}
					}
				}
			case Closed:
				break RunLoop
			}
		}
	}
}
