package engine

import (
	"gopoker/engine/process"
)

func (i *Instance) buildStages() process.Stages {
	return process.Stages{
		process.Stage{
			Name: "rotate-game",
			If: func() bool {
				return i.Gameplay.gameRotation != nil
			},
			Run: i.Gameplay.rotateGame,
		},

		process.Stage{
			Name: "post-antes",
			If:   func() bool { return i.Game.HasAnte || i.Stake.HasAnte() },
			Run:  i.Gameplay.postAntes,
		},

		process.Stage{
			Name: "post-blinds",
			If:   func() bool { return i.Game.HasBlinds },
			Run:  i.Gameplay.postBlinds,
		},

		process.Stage{
			Name: "streets",
			Run:  nil,//i.processStreets,
		},

		process.Stage{
			Name: "showdown",
			Run:  i.showdown,
		},
	}
}
