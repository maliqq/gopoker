package play

import (
	"sync"
)

import (
	"gopoker/play/mode"
	"gopoker/play/street"
)

type State string

const (
	Waiting State = "waiting"
	Active  State = "active"
	Paused  State = "paused"
	Closed  State = "closed"
)

type FSM struct {
	// current state
	State       State
	stateChange chan State
	stateLock   sync.Mutex

	// current mode
	Mode mode.Type

	// current stage
	Stage string

	// current street
	Street street.Type
}

func (this *Play) Start() {
	this.stateChange <- Active
}

func (this *Play) Pause() {
	this.stateChange <- Paused
}

func (this *Play) Resume() {
	this.stateChange <- Active
}

func (this *Play) Close() {
	this.stateChange <- Closed
}
