package engine

import (
  "gopoker/engine/stage"
)

type DealProcess struct {
  g *Gameplay
  stages Stages
}

func CreateDealProcess(g *Gameplay) *DealProcess {
  p := &DealProcess{
    g: g,
    stages: buildStages(g),
  }

  return p
}

func (p *DealProcess) Start() {
  
}

func buildStages(g *Gameplay) Stages {
  return Stages{
    Stage{
      Type: stage.RotateGame,
      If: func() bool { return g.gameRotation != nil },
      Do: g.rotateGame,
    },

    Stage{
      Type: stage.PostAntes,
      If:   func() bool { return g.Game.HasAnte || g.Stake.HasAnte() },
      Do:  g.postAntes,
    },

    Stage{
      Type: stage.PostBlinds,
      If:   func() bool { return g.Game.HasBlinds },
      Do:  g.postBlinds,
    },

    Stage{
      Type: stage.Streets,
      Do:  func() {
        for _, street := range buildStreets(g) {
          street.Start()
        }
      },
    },

    Stage{
      Type: stage.Showdown,
      Do:  g.showdown,
    },
  }
}
