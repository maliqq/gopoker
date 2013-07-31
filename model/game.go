package model

import (
	"fmt"
	"log"
)

import (
	"gopoker/model/game"
	"gopoker/poker/hand"
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

	Hi hand.Ranking
	Lo hand.Ranking

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

func NewGame(g game.Type, limit game.Limit) *Game {
	limitedGame, success := g.(game.LimitedGame)

	if !success {
		log.Printf("got: %#v", g)
		panic("can't create game")
	}

	game := &Game{
		Type:  limitedGame,
		Limit: limit,
	}

	return game.WithDefaults()
}

func (game *Game) WithDefaults() *Game {
	var success bool
	game.GameOptions, success = Games[game.Type]
	if !success {
		log.Printf("got: %#v", game)
		panic("can't populate game with options")
	}

	if game.Limit == "" {
		game.Limit = game.GameOptions.DefaultLimit
	}

	return game
}

func (game *Game) IsMixed() bool {
	return false
}

func (game *Game) String() string {
	return fmt.Sprintf("%s %s", game.Type, game.Limit)
}

func NewMix(g game.Type) *Mix {
	mixedGame, success := g.(game.MixedGame)

	if !success {
		log.Printf("got: %#v\n", g)

		panic("can't create mix")
	}

	mix := &Mix{
		Type: mixedGame,
	}

	return mix.WithDefaults()
}

func (mix *Mix) WithDefaults() *Mix {
	options, success := Mixes[mix.Type]
	if !success {
		log.Printf("got: %#v\n", mix)
		panic("can't populate mix with options")
	}

	games := make([]*Game, len(options))
	for i, mixOptions := range options {
		games[i] = NewGame(mixOptions.Type, mixOptions.Limit)
	}
	mix.Games = games

	return mix
}

func (mix *Mix) String() string {
	return fmt.Sprintf("%s", mix.Type)
}

func (mix *Mix) IsMixed() bool {
	return true
}
