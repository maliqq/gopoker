package gameplay

import (
	"time"
)

import (
	"gopoker/model"
	"gopoker/play/context"
	"gopoker/protocol"
)

type Transition string

const (
	Stop Transition = "stop"
	Next Transition = "next"
)

type GamePlay struct {
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
	Broadcast *protocol.Broadcast `json:"-"`

	Betting             *context.Betting
	*context.Discarding `json:"-"`

	// manage play
	NextDeal chan (<-chan time.Time) `json:"-"`
	Exit     chan int                `json:"-"`
}

func NewGamePlay() *GamePlay {
	return &GamePlay{
		Broadcast: protocol.NewBroadcast(),
		NextDeal:  make(chan (<-chan time.Time)),
		Exit:      make(chan int),
	}
}
