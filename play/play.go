package play

import (
	"fmt"
	"log"
	"time"
)

import (
	"gopoker/exch"
	"gopoker/exch/message"
	"gopoker/model"
	"gopoker/model/bet"
	"gopoker/model/seat"
	"gopoker/play/context"
	"gopoker/play/gameplay"
	"gopoker/play/mode"
	"gopoker/util/console"
)

// Play - play
type Play struct {
	// players action context
	*gameplay.GamePlay

	// finite state machine
	FSM
	// receive protocol messages
	Recv exch.MessageChannel `json:"-"`
}

// NewPlay - create new play
func NewPlay(variation model.Variation, stake *model.Stake) *Play {
	gp := gameplay.NewGamePlay()
	gp.Stake = stake

	play := &Play{
		GamePlay: gp,
		FSM: FSM{
			State:       Waiting,
			stateChange: make(chan State),
			Mode:        mode.Cash, // FIXME
		},
		Recv: make(chan *message.Message),
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
		case msg := <-play.Recv:
			play.handleMessage(msg)

		case newState := <-play.stateChange:
			play.handleStateChange(newState)
		}
	}
}

// Proto - play to protobuf
func (play *Play) Proto() *message.Play {
	return &message.Play{
		Table: play.Table.Proto(),
		Game:  play.Game.Proto(),
		Stake: play.Stake.Proto(),
	}
}
