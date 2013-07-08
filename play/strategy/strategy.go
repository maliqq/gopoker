package strategy

import (
	"log"
)

import (
	"gopoker/play/context"
	"gopoker/play/stage"
)

type stageFunc func(*context.Play)

type Strategy []stageFunc

func (strategy Strategy) Proceed(play *context.Play) {
	for _, context := range strategy {
		context(play)
	}
}

var (
	Default = Strategy{
		stage.Initialize,

		func(play *context.Play) {
			streets, _ := Streets[play.Game.Options.Group]

			for _, street := range streets {
				log.Printf("[play.street] %s\n", street)

				StreetStrategies[street].Proceed(play)
			}
		},

		stage.Finalize,
	}
)
