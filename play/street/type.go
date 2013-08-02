package street

import (
	"gopoker/model/game"
)

// Type - street type
type Type string

// Streets
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

// ByGameGroup - streets by game group
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

// Get - get streets for game group
func Get(group game.Group) []Type {
	return ByGameGroup[group]
}
