package engine

import (
	"github.com/golang/glog"
)

import (
	"gopoker/engine/context"
	"gopoker/engine/stage"
	"gopoker/event"
	"gopoker/message"
)

type DealProcess struct {
	g *Gameplay
	StageStrategy
	Recv event.Channel
	Exit chan bool
}

func NewDealProcess(g *Gameplay) *DealProcess {
	return &DealProcess{
		g:             g,
		StageStrategy: buildStages(g),
		Recv:          make(event.Channel),
		Exit:          make(chan bool),
	}
}

func (process *DealProcess) Run() {
	go process.receive()
	for {
		process.g.d = context.NewDeal()
		process.g.BettingProcess = NewBettingProcess(process.g)

		if !process.StageStrategy.Run() {
			process.Exit <- true
			break
		}
	}
}

func (process *DealProcess) receive() {
DealingLoop:
	for {
		select {
		case <-process.Exit:
			break DealingLoop

		case notification := <-process.Recv:
			glog.Infof("GOT: %#v", notification)
			switch msg := notification.Message.(type) {
			case *message.AddBet:
				process.g.BettingProcess.Recv <- msg.Bet
			}
		}
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
			Do: g.processStreets,
		},

		StageDo{
			Stage: Stage{
				Type: stage.Showdown,
			},
			Do: g.showdown,
		},
	}
}
