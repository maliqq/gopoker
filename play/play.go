package play

import (
	"fmt"
)

import (
	"gopoker/event"
	"gopoker/model"
	"gopoker/play/context"
	"gopoker/play/gameplay"
	"gopoker/play/mode"
	"gopoker/play/street"
)

type State string

// States
const (
	Waiting State = "waiting"
	Active  State = "active"
	Paused  State = "paused"
	Closed  State = "closed"
)

// Play - play
type Play struct {
	// players action context
	*gameplay.Gameplay

	// current state
	State       State
	// current mode
	Mode mode.Type
	// current stage
	Stage string
	// current street
	Street street.Type

	// receive protocol messages
	Call event.Channel `json:"-"`
}

// NewPlay - create new play
func NewPlay(variation model.Variation, stake *model.Stake) *Play {
	gp := gameplay.NewGameplay()
	gp.Stake = stake

	play := &Play{
		Gameplay: gp,
		State:       Waiting,
		Mode:        mode.Cash, // FIXME
		Recv: make(event.Channel),
	}

	// game
	if variation.IsMixed() {
		mix := variation.(*model.Mix)
		play.Mix = mix
		play.GameRotation = context.NewGameRotation(mix, 1)
		play.Game = play.GameRotation.Current()
	} else {
		play.Game = variation.(*model.Game)
	}
	// create table
	play.Table = model.NewTable(play.Game.TableSize)

	go play.receive()

	return play
}

// String - play to string
func (play *Play) String() string {
	return fmt.Sprintf("Game: %s\nTable: %s\n", play.Game, play.Table)
}

func (play *Play) receive() {
	for {
		select {
		case call := <-play.Call:
			switch args := call.(rpc.Call).(type) {
			case rpc.AddBet:
				play.AddBet(...)
			}
		}
	}
}
