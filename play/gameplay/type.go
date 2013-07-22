package gameplay

import (
	"gopoker/model"
	"gopoker/play/command"
	"gopoker/play/context"
	"gopoker/protocol"
)

type Transition string

const (
	Stop Transition = "stop"
	Next Transition = "next"
	Skip Transition = "skip"
	Redo Transition = "redo"
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
	Control chan command.Command `json:"-"`
}
