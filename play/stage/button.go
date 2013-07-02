package stage

import (
	"gopoker/protocol"
)

func (stage *Stage) setButton(pos int) {
	play := stage.Play

	play.Table.SetButton(pos)

	play.Broadcast.All <- protocol.NewMoveButton(pos)
}

func (stage *Stage) moveButton() {
	play := stage.Play

	gameOptions := play.Game.Options

	// button moves only for draw and holdem
	if gameOptions.HasBlinds {
		play.Table.MoveButton()

		play.Broadcast.All <- protocol.NewMoveButton(play.Table.Button)
	}
}
