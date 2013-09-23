package engine

import (
	"log"
)

import (
	"gopoker/engine/stage"
)

type DealProcess struct {
	g *Gameplay
	Stages
}

func CreateDealProcess(g *Gameplay) *DealProcess {
	p := &DealProcess{g: g}
	p.buildStages()

	return p
}

func (p *DealProcess) buildStages() {
	g := p.g
	p.Stages = Stages{
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
