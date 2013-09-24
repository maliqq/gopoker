package engine

func (instance *Instance) Start() {
	instance.stateChange <- InstanceStateChange{NewState: Active}
}

func (instance *Instance) StartAfter(timeout int) {
	instance.stateChange <- InstanceStateChange{Timeout: timeout, NewState: Active}
}

func (instance *Instance) Pause() {
	instance.stateChange <- InstanceStateChange{NewState: Paused}
}

func (instance *Instance) Resume() {
	instance.stateChange <- InstanceStateChange{NewState: Waiting}
}

func (instance *Instance) Stop() {
	instance.stateChange <- InstanceStateChange{NewState: Closed}
}
