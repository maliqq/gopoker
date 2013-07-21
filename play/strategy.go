package play

import (
	"gopoker/model/deal"
	"gopoker/play/mode"
	"gopoker/play/street"
)

type Strategy []Stage

func (strategy Strategy) Proceed(play *Play) {
	for _, stage := range strategy {
		play.Stage = stage.Name
		stage.Invoke(play)
	}
}

var ByStreet = map[street.Type]Strategy{
	// holdem poker
	street.Preflop: Strategy{
		Dealing(deal.Hole, 0),
		Betting,
	},
	street.Flop: Strategy{
		Dealing(deal.Board, 3),
		Betting,
	},
	street.Turn: Strategy{
		Dealing(deal.Board, 1),
		BigBets,
		Betting,
	},
	street.River: Strategy{
		Dealing(deal.Board, 1),
		Betting,
	},

	// seven card poker
	street.Second: Strategy{
		Dealing(deal.Hole, 2),
	},
	street.Third: Strategy{
		Dealing(deal.Door, 1),
		BringIn,
		Betting,
	},
	street.Fourth: Strategy{
		Dealing(deal.Door, 1),
		Betting,
	},
	street.Fifth: Strategy{
		Dealing(deal.Door, 1),
		BigBets,
		Betting,
	},
	street.Sixth: Strategy{
		Dealing(deal.Door, 1),
		Betting,
	},
	street.Seventh: Strategy{
		Dealing(deal.Hole, 1),
		Betting,
	},

	// draw poker
	street.Predraw: Strategy{
		Dealing(deal.Hole, 5),
		Betting,
		Discarding,
	},
	street.Draw: Strategy{
		BigBets,
		Betting,
		Discarding,
	},
	street.FirstDraw: Strategy{
		Betting,
		Discarding,
	},
	street.SecondDraw: Strategy{
		Betting,
		Discarding,
	},
	street.ThirdDraw: Strategy{
		Betting,
		Discarding,
	},
}

var ByMode = map[mode.Type]Strategy{
	mode.Cash: Strategy{
		DealStart,
		PostAntes,
		PostBlinds,
		Streets,
		Showdown,
		DealStop,
	},
}
