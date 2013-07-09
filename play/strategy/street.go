package strategy

import (
	"gopoker/model/deal"
	"gopoker/model/game"
	"gopoker/play/stage"
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
		stage.Dealing(deal.Hole, 0),
		stage.BettingRound,
	},
	Flop: Strategy{
		stage.Dealing(deal.Board, 3),
		stage.BettingRound,
	},
	Turn: Strategy{
		stage.Dealing(deal.Board, 1),
		stage.BigBets,
		stage.BettingRound,
	},
	River: Strategy{
		stage.Dealing(deal.Board, 1),
		stage.BettingRound,
	},

	// seven card poker
	Second: Strategy{
		stage.Dealing(deal.Hole, 2),
	},
	Third: Strategy{
		stage.Dealing(deal.Door, 1),
		stage.BringIn,
		stage.BettingRound,
	},
	Fourth: Strategy{
		stage.Dealing(deal.Door, 1),
		stage.BettingRound,
	},
	Fifth: Strategy{
		stage.Dealing(deal.Door, 1),
		stage.BigBets,
		stage.BettingRound,
	},
	Sixth: Strategy{
		stage.Dealing(deal.Door, 1),
		stage.BettingRound,
	},
	Seventh: Strategy{
		stage.Dealing(deal.Hole, 1),
		stage.BettingRound,
	},

	// draw poker
	Predraw: Strategy{
		stage.Dealing(deal.Hole, 5),
		stage.BettingRound,
		stage.DiscardingRound,
	},
	Draw: Strategy{
		stage.BigBets,
		stage.BettingRound,
		stage.DiscardingRound,
	},
	FirstDraw: Strategy{
		stage.BettingRound,
		stage.DiscardingRound,
	},
	SecondDraw: Strategy{
		stage.BettingRound,
		stage.DiscardingRound,
	},
	ThirdDraw: Strategy{
		stage.BettingRound,
		stage.DiscardingRound,
	},
}
