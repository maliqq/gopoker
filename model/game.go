package model

import (
	"fmt"
)

import (
	"gopoker/model/game"
	"gopoker/poker/ranking"
)

type GameOptions struct {
	Group game.Group

	HasBlinds  bool
	HasAnte    bool
	HasBringIn bool
	HasBoard   bool
	HasVela    bool

	Discards  bool
	Reshuffle bool
	Max       int
	Pocket    int
	Streets   int

	Hi ranking.Type
	Lo ranking.Type

	DefaultLimit game.Limit
}

type MixOptions struct {
	game.Type
	game.Limit
}

type Game struct {
	game.Type
	game.Limit
	Stake   *game.Stake
	Options *GameOptions
}

type Mix struct {
	game.Type
	*game.Stake
	Games []*Game
}

type Variation interface {
}

var Games = map[game.LimitedGame]*GameOptions{
	game.Texas: &GameOptions{
		Group:        game.Holdem,
		HasBoard:     true,
		HasBlinds:    true,
		Max:          10,
		Hi:           ranking.High,
		Pocket:       2,
		DefaultLimit: game.NoLimit,
	},

	game.Omaha: &GameOptions{
		Group:        game.Holdem,
		HasBoard:     true,
		HasBlinds:    true,
		Max:          10,
		Pocket:       4,
		Hi:           ranking.High,
		DefaultLimit: game.PotLimit,
	},

	game.Omaha8: &GameOptions{
		Group:        game.Holdem,
		HasBoard:     true,
		HasBlinds:    true,
		Max:          10,
		Pocket:       4,
		Hi:           ranking.High,
		Lo:           ranking.AceFive8,
		DefaultLimit: game.PotLimit,
	},

	game.Stud: &GameOptions{
		Group:        game.SevenCard,
		HasAnte:      true,
		HasBringIn:   true,
		HasVela:      true,
		Max:          8,
		Pocket:       7,
		Hi:           ranking.High,
		DefaultLimit: game.FixedLimit,
	},

	game.Stud8: &GameOptions{
		Group:        game.SevenCard,
		HasAnte:      true,
		HasBringIn:   true,
		HasVela:      true,
		Max:          8,
		Pocket:       7,
		Hi:           ranking.High,
		Lo:           ranking.AceFive8,
		DefaultLimit: game.FixedLimit,
	},

	game.Razz: &GameOptions{
		Group:        game.SevenCard,
		HasAnte:      true,
		HasBringIn:   true,
		HasVela:      true,
		Max:          8,
		Pocket:       7,
		Hi:           ranking.AceFive,
		DefaultLimit: game.FixedLimit,
	},

	game.London: &GameOptions{
		Group:        game.SevenCard,
		HasAnte:      true,
		HasBringIn:   true,
		HasVela:      true,
		Max:          8,
		Pocket:       7,
		Hi:           ranking.AceSix,
		DefaultLimit: game.FixedLimit,
	},

	game.FiveCard: &GameOptions{
		Group:        game.SingleDraw,
		HasBlinds:    true,
		Discards:     true,
		Reshuffle:    true,
		Max:          6,
		Pocket:       5,
		Streets:      1,
		Hi:           ranking.High,
		DefaultLimit: game.FixedLimit,
	},

	game.Single27: &GameOptions{
		Group:        game.SingleDraw,
		HasBlinds:    true,
		Discards:     true,
		Reshuffle:    true,
		Max:          6,
		Pocket:       5,
		Streets:      1,
		Hi:           ranking.DeuceSeven,
		DefaultLimit: game.FixedLimit,
	},

	game.Triple27: &GameOptions{
		Group:        game.TripleDraw,
		HasBlinds:    true,
		Discards:     true,
		Reshuffle:    true,
		Max:          6,
		Pocket:       5,
		Streets:      3,
		Hi:           ranking.DeuceSeven,
		DefaultLimit: game.FixedLimit,
	},

	game.Badugi: &GameOptions{
		Group:        game.TripleDraw,
		HasBlinds:    true,
		Discards:     true,
		Reshuffle:    true,
		Max:          6,
		Pocket:       4,
		Hi:           ranking.Badugi,
		DefaultLimit: game.FixedLimit,
	},
}

var Mixes = map[game.MixedGame][]MixOptions{
	game.Horse: []MixOptions{
		MixOptions{Type: game.Texas, Limit: game.FixedLimit},
		MixOptions{Type: game.Omaha8, Limit: game.FixedLimit},
		MixOptions{Type: game.Razz, Limit: game.FixedLimit},
		MixOptions{Type: game.Stud, Limit: game.FixedLimit},
		MixOptions{Type: game.Stud8, Limit: game.FixedLimit},
	},

	game.Eight: []MixOptions{
		MixOptions{Type: game.Triple27, Limit: game.FixedLimit},
		MixOptions{Type: game.Texas, Limit: game.FixedLimit},
		MixOptions{Type: game.Omaha8, Limit: game.FixedLimit},
		MixOptions{Type: game.Razz, Limit: game.FixedLimit},
		MixOptions{Type: game.Stud, Limit: game.FixedLimit},
		MixOptions{Type: game.Stud8, Limit: game.FixedLimit},
		MixOptions{Type: game.Texas, Limit: game.NoLimit},
		MixOptions{Type: game.Omaha, Limit: game.PotLimit},
	},
}

func NewGame(variation game.Type, limit game.Limit, stake *game.Stake) *Game {
	switch variation.(type) {
	case game.LimitedGame:
		options, _ := Games[variation.(game.LimitedGame)]

		return &Game{
			Type:    variation,
			Limit:   limit,
			Stake:   stake,
			Options: options,
		}
	
	default:
		fmt.Printf("got: %s\n", variation)
		panic("can't create game")
	}
}

func (game *Game) String() string {
	return fmt.Sprintf("%s %s %s", game.Type, game.Limit, game.Stake)
}

func NewMix(variation game.Type, stake *game.Stake) *Mix {
	switch variation.(type) {
	case game.MixedGame:
		options, _ := Mixes[variation.(game.MixedGame)]

		games := make([]*Game, len(options))
		for i, mixOptions := range options {
			games[i] = NewGame(mixOptions.Type, mixOptions.Limit, stake)
		}

		return &Mix{
			Type:  variation,
			Stake: stake,
			Games: games,
		}
	
	default:
		fmt.Printf("got: %s\n", variation)
		panic("can't create mix")
	}
}

func (mix *Mix) String() string {
	return fmt.Sprintf("%s %s", mix.Type, mix.Stake)
}
