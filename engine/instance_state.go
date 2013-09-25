package engine

import (
	"gopoker/engine/state"
)

func (instance *Instance) Start() {
	instance.stateChange <- InstanceStateChange{NewState: state.Active}
}

func (instance *Instance) StartAfter(timeout int) {
	instance.stateChange <- InstanceStateChange{Timeout: timeout, NewState: state.Active}
}

func (instance *Instance) Pause() {
	instance.stateChange <- InstanceStateChange{NewState: state.Paused}
}

func (instance *Instance) Resume() {
	instance.stateChange <- InstanceStateChange{NewState: state.Waiting}
}

func (instance *Instance) Stop() {
	instance.stateChange <- InstanceStateChange{NewState: state.Closed}
}
