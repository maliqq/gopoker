package engine

import (
	"log"
)

import (
	"gopoker/engine/context"
	"gopoker/engine/stage"
)

type DealProcess struct {
	StageStrategy
}

func NewDealProcess(g *Gameplay) *DealProcess {
	g.d = context.NewDeal()

	return &DealProcess{
		StageStrategy: buildStages(g),
	}
}

func buildStages(g *Gameplay) StageStrategy {
	return StageStrategy{
		StageDo{
			Stage: Stage{
				Type: stage.PrepareSeats,
			},
			Do: g.prepareSeats,
		},

		StageDo{
			Stage: Stage{
				Type: stage.RotateGame,
				If:   func() bool { return g.gameRotation != nil },
			},
			Do: g.rotateGame,
		},

		StageDo{
			Stage: Stage{
				Type: stage.PostAntes,
				If:   func() bool { return g.Game.HasAnte || g.Stake.HasAnte() },
			},
			Do: g.postAntes,
		},

		StageExit{
			Stage: Stage{
				Type: stage.PostBlinds,
				If:   func() bool { return g.Game.HasBlinds },
			},
			Do: g.postBlinds,
		},

		StageSkip{
			Stage: Stage{
				Type: stage.Streets,
			},
			Do: func(skip chan bool) {
				for _, street := range buildStreets(g) {
					log.Printf("[street] %s", street.Type)
					if !street.Run() {
						skip <- true
						return
					}
				}
				close(skip)
			},
		},

		StageDo{
			Stage: Stage{
				Type: stage.Showdown,
			},
			Do: g.showdown,
		},
	}
}
