package engine

func (i *Instance) buildStages() StageProcesses {
	return StageProcesses{
		StageProcess{
			Name: "start-deal",
			Run:  i.startDeal,
		},

		StageProcess{
			Name: "rotate-game",
			If: func() bool {
				return i.Gameplay.gameRotation != nil
			},
			Run: i.Gameplay.rotateGame,
		},

		StageProcess{
			Name: "post-antes",
			If:   func() bool { return i.Game.HasAnte || i.Stake.HasAnte() },
			Run:  i.Gameplay.postAntes,
		},

		StageProcess{
			Name: "post-blinds",
			If:   func() bool { return i.Game.HasBlinds },
			Run:  i.Gameplay.postBlinds,
		},

		StageProcess{
			Name: "streets",
			Run:  i.processStreets,
		},

		StageProcess{
			Name: "showdown",
			Run:  i.showdown,
		},

		StageProcess{
			Name: "deal-stop",
			Run:  i.stopDeal,
		},
	}
}
