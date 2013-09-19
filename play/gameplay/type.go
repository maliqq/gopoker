package gameplay

import (
	"time"
)

import (
	"gopoker/event"
	"gopoker/model"
	"gopoker/play/context"
)

// Transition for stage
type Transition string

// Transitions
const (
	Stop Transition = "stop"
	Next Transition = "next"
)

// Gameplay - gameplay context
type Gameplay struct {
	// dealt cards context
	Deal *model.Deal

	// mixed or limited game
	Game                  *model.Game
	Mix                   *model.Mix
	*context.GameRotation `json:"-"`

	// betting price
	Stake *model.Stake

	// players seating context
	Table *model.Table

	// broadcast protocol messages
	Broadcast *event.Broadcast `json:"-"`

	Betting             *context.Betting
	*context.Discarding `json:"-"`

	// manage play
	NextDeal chan (<-chan time.Time) `json:"-"`
	Exit     chan int                `json:"-"`
}

// NewGameplay - create gameplay context
func NewGameplay() *Gameplay {
	return &Gameplay{
		Broadcast: event.NewBroadcast(),
		NextDeal:  make(chan (<-chan time.Time)),
		Exit:      make(chan int),
	}
}
