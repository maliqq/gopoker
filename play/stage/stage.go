package stage

import (
	"gopoker/play/context"
)

type Stage struct {
	*context.Play
	*context.Betting
}

func NewStage(play *context.Play) *Stage {
	return &Stage{
		Play:    play,
		Betting: context.NewBetting(),
	}
}
