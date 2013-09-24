package engine

import (
	"gopoker/message"
	"gopoker/model/bet"
)

func (g *Gameplay) postAntes() {
	active := g.Table.Ring().Active()
	g.b.NewRound(active)

	for _, box := range active {
		newBet := g.b.ForceBet(bet.Ante, g.Stake)

		g.b.AddBet(newBet)

		g.e.Notify(
			&message.AddBet{box.Pos, newBet},
		).All()
		g.b.Round.Move()
	}

	g.completeBetting()
}
