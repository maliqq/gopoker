package play

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
		BettingRound,
	},
	Flop: Strategy{
		Dealing(deal.Board, 3),
		BettingRound,
	},
	Turn: Strategy{
		Dealing(deal.Board, 1),
		BigBets,
		BettingRound,
	},
	River: Strategy{
		Dealing(deal.Board, 1),
		BettingRound,
	},

	// seven card poker
	Second: Strategy{
		Dealing(deal.Hole, 2),
	},
	Third: Strategy{
		Dealing(deal.Door, 1),
		BringIn,
		BettingRound,
	},
	Fourth: Strategy{
		Dealing(deal.Door, 1),
		BettingRound,
	},
	Fifth: Strategy{
		Dealing(deal.Door, 1),
		BigBets,
		BettingRound,
	},
	Sixth: Strategy{
		Dealing(deal.Door, 1),
		BettingRound,
	},
	Seventh: Strategy{
		Dealing(deal.Hole, 1),
		BettingRound,
	},

	// draw poker
	Predraw: Strategy{
		Dealing(deal.Hole, 5),
		BettingRound,
		DiscardingRound,
	},
	Draw: Strategy{
		BigBets,
		BettingRound,
		DiscardingRound,
	},
	FirstDraw: Strategy{
		BettingRound,
		DiscardingRound,
	},
	SecondDraw: Strategy{
		BettingRound,
		DiscardingRound,
	},
	ThirdDraw: Strategy{
		BettingRound,
		DiscardingRound,
	},
}
