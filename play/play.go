package play

import (
	"gopoker/model"
	"gopoker/play/context"
	"gopoker/play/stage"
)

func Start(play *context.Play) {
	play.Deal = model.NewDeal()

	stage := stage.NewStage(play)

	DefaultStrategy.Proceed(stage)
}
