package engine

import (
	"gopoker/engine/context"
	"gopoker/model"
	"gopoker/model/seat"
	"gopoker/util"
)

type Context struct {
	Stake        *model.Stake
	Table        *model.Table
	Game         *model.Game
	Mix          *model.Mix
	gameRotation *util.Rotation
}

type Gameplay struct {
	*Context

	d           *context.Deal
	DealProcess *DealProcess

	b              *context.Betting
	BettingProcess *BettingProcess

	e *context.Broker
}

func NewGameplay(ctx *Context) *Gameplay {
	g := &Gameplay{
		Context: ctx,
		e:       context.NewBroker(),
	}

	// FIXME
	if g.Mix != nil {
		g.gameRotation = util.NewRotation(g.Mix, 0)
		g.setCurrentGame()
	}

	return g
}

func (g *Gameplay) Betting() *context.Betting {
	return g.b
}

func (g *Gameplay) Broker() *context.Broker {
	return g.e
}

func (g *Gameplay) Deal() *context.Deal {
	return g.d
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

func (g *Gameplay) prepareSeats() {
	for _, s := range g.Table.Seats {
		switch s.State {
		case seat.Ready, seat.Play, seat.Fold:
			s.Play()
		}
	}
}
