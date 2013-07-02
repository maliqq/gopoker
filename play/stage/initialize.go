package stage

import (
	"log"
)

import (
	"gopoker/model/seat"
)

func (stage *Stage) Initialize() {
	play := stage.Play

	gameOptions := play.Game.Options
	stake := play.Game.Stake

	// reset seats
	log.Printf("[play.start] reset seats\n")

	for _, s := range play.Table.Seats {
		switch s.State {
		case seat.Ready, seat.Play:
			s.SetPlaying()
		}
	}

	if gameOptions.HasAnte || stake.HasAnte() {
		log.Printf("[play.start] post antes\n")

		stage.postAntes()
	}

	if gameOptions.HasBlinds {
		log.Printf("[play.start] post blinds\n")

		stage.postBlinds()
	}
}
