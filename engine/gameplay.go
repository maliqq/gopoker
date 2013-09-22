package engine

import (
	"gopoker/hub"
	_"gopoker/message"
	"gopoker/model"
	"gopoker/model/seat"
	"gopoker/util"
)

type Context struct {
	Game  *model.Game
	Stake *model.Stake
	Mix   *model.Mix
	Table *model.Table
}

type Gameplay struct {
	Deal *model.Deal
	
	gameRotation *util.Rotation
	*Context

	b *BettingContext
	e *hub.Broker
}

func NewGameplay(context *Context) *Gameplay {
	g := &Gameplay{
		Context: context,
		e: hub.NewBroker(),
	}

	if g.Mix != nil {
		g.gameRotation = util.NewRotation(g.Mix, 0)
		g.setCurrentGame()
	}

	return g
}

func (g *Gameplay) setCurrentGame() {
	g.Game = g.Mix.Games[g.gameRotation.Current()]
}

func (g *Gameplay) rotateGame() {
	if g.gameRotation.Move() {
		g.setCurrentGame()
	}
}

func (g *Gameplay) turnOnBigBets() {
	g.b.BigBets()
}

func (g *Gameplay) startDeal() {
	for _, s := range g.Table.Seats {
		switch s.State {
		case seat.Ready, seat.Play, seat.Fold:
			s.Play()
		}
	}

	g.Deal = model.NewDeal()

	g.b = NewBettingContext()
}

func (i *Instance) stopDeal() {
}
