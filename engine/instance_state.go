package engine

func (instance *Instance) Start() {
	instance.StateChange <- InstanceStateChange{NewState: Active}
}

func (instance *Instance) StartAfter(timeout int) {
	instance.StateChange <- InstanceStateChange{Timeout: timeout, NewState: Active}
}

func (instance *Instance) Pause() {
	instance.StateChange <- InstanceStateChange{NewState: Paused}
}

func (instance *Instance) Resume() {
	instance.StateChange <- InstanceStateChange{NewState: Waiting}
}

func (instance *Instance) Stop() {
	instance.StateChange <- InstanceStateChange{NewState: Closed}
}
