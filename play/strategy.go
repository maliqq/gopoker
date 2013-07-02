package play

import (
	"log"
)

import (
	"gopoker/play/stage"
)

type Strategy []Stage

func (strategy Strategy) Proceed(stage *stage.Stage) {
	for _, context := range strategy {
		context(stage)
	}
}

var (
	DefaultStrategy = Strategy{
		// start
		func(stage *stage.Stage) {
			log.Printf("[play] start\n")

			stage.Start()
		},

		// streets
		func(stage *stage.Stage) {
			game := stage.Play.Game

			streets, _ := Streets[game.Options.Group]
			for _, street := range streets {
				log.Printf("[play.street] %s\n", street)

				StreetStrategies[street].Proceed(stage)
			}
		},

		// stop
		func(stage *stage.Stage) {
			log.Printf("[play] stop")

			stage.Stop()
		},
	}
)
