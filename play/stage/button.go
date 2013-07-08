package stage

import (
	"gopoker/protocol"
	"gopoker/play/context"
)

func setButton(play *context.Play, pos int) {
	play.Table.SetButton(pos)

	play.Broadcast.All <- protocol.NewMoveButton(pos)
}

func moveButton(play *context.Play) {
	gameOptions := play.Game.Options

	// button moves only for draw and holdem
	if gameOptions.HasBlinds {
		play.Table.MoveButton()

		play.Broadcast.All <- protocol.NewMoveButton(play.Table.Button)
	}
}
