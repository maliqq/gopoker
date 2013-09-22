package engine

import (
	"gopoker/hub"
	"gopoker/model"
	"gopoker/model/seat"
  "gopoker/engine/context"
	"gopoker/util"
)

type Context struct {
  Game  *model.Game
  Stake *model.Stake
  Mix   *model.Mix
  Table *model.Table
}

type Gameplay struct {
	*Context
  gameRotation *util.Rotation

	d *context.Deal
	b *context.Betting

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

func (g *Gameplay) prepareSeats() {
  for _, s := range g.Table.Seats {
    switch s.State {
    case seat.Ready, seat.Play, seat.Fold:
      s.Play()
    }
  }
}
