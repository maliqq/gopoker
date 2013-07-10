package street

import (
	"gopoker/model/game"
)

type Type string

const (
	Preflop Type = "preflop"
	Flop    Type = "flop"
	Turn    Type = "turn"
	River   Type = "river"

	Second  Type = "second"
	Third   Type = "third"
	Fourth  Type = "fourth"
	Fifth   Type = "fifth"
	Sixth   Type = "sixth"
	Seventh Type = "seventh"

	Predraw    Type = "predraw"
	Draw       Type = "draw"
	FirstDraw  Type = "first-draw"
	SecondDraw Type = "second-draw"
	ThirdDraw  Type = "third-draw"
)

var ByGameGroup = map[game.Group][]Type{
	game.Holdem: []Type{
		Preflop, Flop, Turn, River,
	},

	game.SevenCard: []Type{
		Second, Third, Fourth, Fifth, Seventh,
	},

	game.SingleDraw: []Type{
		Predraw, Draw,
	},

	game.TripleDraw: []Type{
		Predraw, FirstDraw, SecondDraw, ThirdDraw,
	},
}

func Get(group game.Group) []Type {
	return ByGameGroup[group]
}
