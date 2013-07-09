package stage

import (
	"log"
)

import (
	"gopoker/model/seat"
	"gopoker/play/context"
)

func Initialize(play *context.Play) {
	log.Println("[play] initialize")

	gameOptions := play.Game.Options
	stake := play.Game.Stake

	// reset seats
	log.Printf("[play.initialize] reset seats\n")

	for _, s := range play.Table.Seats {
		switch s.State {
		case seat.Ready, seat.Play:
			s.SetPlaying()
		}
	}

	if gameOptions.HasAnte || stake.HasAnte() {
		log.Printf("[play.initialize] post antes\n")

		postAntes(play)
	}

	if gameOptions.HasBlinds {
		log.Printf("[play.initialize] post blinds\n")

		postBlinds(play)
	}
}
