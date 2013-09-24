package engine

import (
	"fmt"
	"time"
)

import (
	"gopoker/engine/mode"
	"gopoker/engine/street"
	"gopoker/event"
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

	Recv event.Channel

	State       State
	stateChange chan InstanceStateChange
	Street      street.Type
	Mode        mode.Type
}

func NewInstance(context *Context) *Instance {
	instance := &Instance{
		State:       Waiting,
		stateChange: make(chan InstanceStateChange),
		Recv:        make(event.Channel),
		Gameplay:    NewGameplay(context),
	}

	go instance.receive()

	return instance
}

func (instance *Instance) doStart() {
	fmt.Println("start...")
	instance.State = Active

	instance.BettingProcess = NewBettingProcess(instance.Gameplay)
	instance.DealProcess = NewDealProcess(instance.Gameplay)
	instance.DealProcess.Run()
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
		if event.NewState != Active && event.NewState != Closed {
			return false
		}
	}

	if timeout := event.Timeout; timeout != 0 {
		<-time.After(time.Duration(timeout) * time.Second)
	}
	instance.State = event.NewState

	return true
}

func (instance *Instance) receive() {
RunLoop:
	for {
		select {
		case <-instance.Recv:
		case event := <-instance.stateChange:
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
					case event := <-instance.stateChange:
						if instance.processStateChange(event) {
							break PauseLoop
						}
					}
				}
			case Closed:
				break RunLoop
			}
		default:
		}

	}
}
