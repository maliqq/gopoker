package strategy

import (
	"gopoker/model/deal"
	"gopoker/model/game"
)

type Street string

const (
	Preflop Street = "preflop"
	Flop    Street = "flop"
	Turn    Street = "turn"
	River   Street = "river"

	Second  Street = "second"
	Third   Street = "third"
	Fourth  Street = "fourth"
	Fifth   Street = "fifth"
	Sixth   Street = "sixth"
	Seventh Street = "seventh"

	Predraw    Street = "predraw"
	Draw       Street = "draw"
	FirstDraw  Street = "first-draw"
	SecondDraw Street = "second-draw"
	ThirdDraw  Street = "third-draw"
)

var Streets = map[game.Group][]Street{
	game.Holdem: []Street{
		Preflop, Flop, Turn, River,
	},

	game.SevenCard: []Street{
		Second, Third, Fourth, Fifth, Seventh,
	},

	game.SingleDraw: []Street{
		Predraw, Draw,
	},

	game.TripleDraw: []Street{
		Predraw, FirstDraw, SecondDraw, ThirdDraw,
	},
}

var StreetStrategies = map[Street]Strategy{
	// holdem poker
	Preflop: Strategy{
		Dealing(deal.Hole, 0),
		Betting,
	},
	Flop: Strategy{
		Dealing(deal.Board, 3),
		Betting,
	},
	Turn: Strategy{
		Dealing(deal.Board, 1),
		BigBets,
		Betting,
	},
	River: Strategy{
		Dealing(deal.Board, 1),
		Betting,
	},

	// seven card poker
	Second: Strategy{
		Dealing(deal.Hole, 2),
	},
	Third: Strategy{
		Dealing(deal.Door, 1),
		BringIn,
		Betting,
	},
	Fourth: Strategy{
		Dealing(deal.Door, 1),
		Betting,
	},
	Fifth: Strategy{
		Dealing(deal.Door, 1),
		BigBets,
		Betting,
	},
	Sixth: Strategy{
		Dealing(deal.Door, 1),
		Betting,
	},
	Seventh: Strategy{
		Dealing(deal.Hole, 1),
		Betting,
	},

	// draw poker
	Predraw: Strategy{
		Dealing(deal.Hole, 5),
		Betting,
		Discarding,
	},
	Draw: Strategy{
		BigBets,
		Betting,
		Discarding,
	},
	FirstDraw: Strategy{
		Betting,
		Discarding,
	},
	SecondDraw: Strategy{
		Betting,
		Discarding,
	},
	ThirdDraw: Strategy{
		Betting,
		Discarding,
	},
}
