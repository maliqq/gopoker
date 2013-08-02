package play

import (
	"sync"
)

import (
	"gopoker/play/mode"
	"gopoker/play/street"
)

// State - play state
type State string

// States
const (
	Waiting State = "waiting"
	Active  State = "active"
	Paused  State = "paused"
	Closed  State = "closed"
)

// FSM - finite state machine
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

// Start - start play
func (play *Play) Start() {
	play.stateChange <- Active
}

// Pause - pause play
func (play *Play) Pause() {
	play.stateChange <- Paused
}

// Resume - resume play
func (play *Play) Resume() {
	play.stateChange <- Active
}

// Close - close play
func (play *Play) Close() {
	play.stateChange <- Closed
}
