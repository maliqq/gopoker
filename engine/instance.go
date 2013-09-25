package engine

import (
	"time"
)

import (
	"github.com/golang/glog"
)

import (
	"gopoker/engine/state"
)

type InstanceStateChange struct {
	Timeout  int
	NewState state.Type
}

type Instance struct {
	*Gameplay

	State       state.Type
	stateChange chan InstanceStateChange
}

func NewInstance(context *Context) *Instance {
	instance := &Instance{
		State:       state.Waiting,
		stateChange: make(chan InstanceStateChange),
		Gameplay:    NewGameplay(context),
	}

	go instance.receive()

	return instance
}

func (instance *Instance) doStart() {
	glog.Info("start...")
	instance.State = state.Active

	instance.DealProcess = NewDealProcess(instance.Gameplay)
	instance.DealProcess.Run()
}

func (instance *Instance) doPause() {
	glog.Info("pause...")
	instance.State = state.Paused
}

func (instance *Instance) doResume() {
	glog.Info("resume...")
	instance.State = state.Waiting
}

func (instance *Instance) doStop() {
	glog.Info("stop...")
	instance.State = state.Closed
}

func (instance *Instance) processStateChange(event InstanceStateChange) bool {
	// check for correct transition
	switch instance.State {
	case state.Paused, state.Closed:
		if event.NewState != state.Waiting {
			return false
		}
	case state.Waiting:
		if event.NewState != state.Active && event.NewState != state.Closed {
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
		//case <-instance.Recv:
		case event := <-instance.stateChange:
			if !instance.processStateChange(event) {
				break
			}

			switch event.NewState {
			
			case state.Active:
				instance.doStart()
			
			case state.Paused:
			PauseLoop:
				for {
					select {
					case event := <-instance.stateChange:
						if instance.processStateChange(event) {
							break PauseLoop
						}
					}
				}
			
			case state.Closed:
				break RunLoop
			}
		default:
		}

	}
}
