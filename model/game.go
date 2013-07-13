package model

import (
	"fmt"
)

import (
	"gopoker/model/game"
	"gopoker/poker/ranking"
)

const (
	MixedGameMaxTableSize = 8
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

	MaxTableSize int
	PocketSize   int
	Streets      int // number of streets till showdown

	Hi ranking.Type
	Lo ranking.Type

	DefaultLimit game.Limit
}

type MixOptions struct {
	Type game.LimitedGame
	game.Limit
}

type Game struct {
	Type game.LimitedGame
	game.Limit
	*GameOptions `json:"-"`
}

type Mix struct {
	Type  game.MixedGame
	Games []*Game `json:"-"`
}

type Variation interface {
	IsMixed() bool
}

var Games = map[game.LimitedGame]*GameOptions{
	game.Texas: &GameOptions{
		Group:        game.Holdem,
		HasBoard:     true,
		HasBlinds:    true,
		MaxTableSize: 10,
		Hi:           ranking.High,
		PocketSize:   2,
		DefaultLimit: game.NoLimit,
	},

	game.Omaha: &GameOptions{
		Group:        game.Holdem,
		HasBoard:     true,
		HasBlinds:    true,
		MaxTableSize: 10,
		PocketSize:   4,
		Hi:           ranking.High,
		DefaultLimit: game.PotLimit,
	},

	game.Omaha8: &GameOptions{
		Group:        game.Holdem,
		HasBoard:     true,
		HasBlinds:    true,
		MaxTableSize: 10,
		PocketSize:   4,
		Hi:           ranking.High,
		Lo:           ranking.AceFive8,
		DefaultLimit: game.PotLimit,
	},

	game.Stud: &GameOptions{
		Group:        game.SevenCard,
		HasAnte:      true,
		HasBringIn:   true,
		HasVela:      true,
		MaxTableSize: 8,
		PocketSize:   7,
		Hi:           ranking.High,
		DefaultLimit: game.FixedLimit,
	},

	game.Stud8: &GameOptions{
		Group:        game.SevenCard,
		HasAnte:      true,
		HasBringIn:   true,
		HasVela:      true,
		MaxTableSize: 8,
		PocketSize:   7,
		Hi:           ranking.High,
		Lo:           ranking.AceFive8,
		DefaultLimit: game.FixedLimit,
	},

	game.Razz: &GameOptions{
		Group:        game.SevenCard,
		HasAnte:      true,
		HasBringIn:   true,
		HasVela:      true,
		MaxTableSize: 8,
		PocketSize:   7,
		Hi:           ranking.AceFive,
		DefaultLimit: game.FixedLimit,
	},

	game.London: &GameOptions{
		Group:        game.SevenCard,
		HasAnte:      true,
		HasBringIn:   true,
		HasVela:      true,
		MaxTableSize: 8,
		PocketSize:   7,
		Hi:           ranking.AceSix,
		DefaultLimit: game.FixedLimit,
	},

	game.FiveCard: &GameOptions{
		Group:        game.SingleDraw,
		HasBlinds:    true,
		Discards:     true,
		Reshuffle:    true,
		MaxTableSize: 6,
		PocketSize:   5,
		Streets:      1,
		Hi:           ranking.High,
		DefaultLimit: game.FixedLimit,
	},

	game.Single27: &GameOptions{
		Group:        game.SingleDraw,
		HasBlinds:    true,
		Discards:     true,
		Reshuffle:    true,
		MaxTableSize: 6,
		PocketSize:   5,
		Streets:      1,
		Hi:           ranking.DeuceSeven,
		DefaultLimit: game.FixedLimit,
	},

	game.Triple27: &GameOptions{
		Group:        game.TripleDraw,
		HasBlinds:    true,
		Discards:     true,
		Reshuffle:    true,
		MaxTableSize: 6,
		PocketSize:   5,
		Streets:      3,
		Hi:           ranking.DeuceSeven,
		DefaultLimit: game.FixedLimit,
	},

	game.Badugi: &GameOptions{
		Group:        game.TripleDraw,
		HasBlinds:    true,
		Discards:     true,
		Reshuffle:    true,
		MaxTableSize: 6,
		PocketSize:   4,
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

func NewGame(g game.Type, limit game.Limit) *Game {
	limitedGame, success := g.(game.LimitedGame)

	if !success {
		fmt.Printf("got: %s\n", g)
		panic("can't create game")
	}

	options := Games[limitedGame]

	if limit == "" {
		limit = options.DefaultLimit
	}

	return &Game{
		Type:        limitedGame,
		Limit:       limit,
		GameOptions: options,
	}
}

func (game *Game) IsMixed() bool {
	return false
}

func (game *Game) String() string {
	return fmt.Sprintf("%s %s %s", game.Type, game.Limit)
}

func NewMix(g game.Type) *Mix {
	mixedGame, success := g.(game.MixedGame)

	if !success {
		fmt.Printf("got: %s\n", g)

		panic("can't create mix")
	}

	options, _ := Mixes[mixedGame]

	games := make([]*Game, len(options))
	for i, mixOptions := range options {
		games[i] = NewGame(mixOptions.Type, mixOptions.Limit)
	}

	return &Mix{
		Type:  mixedGame,
		Games: games,
	}
}

func (mix *Mix) String() string {
	return fmt.Sprintf("%s %s", mix.Type)
}

func (mix *Mix) IsMixed() bool {
	return true
}
