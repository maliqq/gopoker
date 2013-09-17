package play

import (
	"fmt"
)

import (
	"gopoker/event"
	"gopoker/event/message/format/protobuf"
	"gopoker/model"
	"gopoker/play/context"
	"gopoker/play/gameplay"
	"gopoker/play/mode"
)

// Play - play
type Play struct {
	// players action context
	*gameplay.GamePlay

	// finite state machine
	FSM
	// receive protocol messages
	Recv event.Channel `json:"-"`
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
		case event := <-play.Recv:
			play.HandleEvent(event)

		case newState := <-play.stateChange:
			play.HandleStateChange(newState)
		}
	}
}

// Proto - play to protobuf
func (play *Play) ProtoPlay() *protobuf.Play {
	return &protobuf.Play{
		Table: play.Table.Proto(),
		Game:  play.Game.Proto(),
		Stake: play.Stake.Proto(),
	}
}
