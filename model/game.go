package model

import (
	"fmt"
	"log"
)

import (
	"gopoker/model/game"
	"gopoker/poker/hand"
	"gopoker/protocol/message"
)

const (
	// MixedGameMaxTableSize - table size for mixed game
	MixedGameMaxTableSize = 8
)

// GameOptions - game options
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

// MixOptions - mix options
type MixOptions struct {
	Type game.LimitedGame
	game.Limit
}

// Game - game
type Game struct {
	Type game.LimitedGame
	game.Limit
	*GameOptions `json:"-"`
}

// Mix - mix
type Mix struct {
	Type  game.MixedGame
	Games []*Game `json:"-"`
}

// Variation - union of Game and Mix
type Variation interface {
	IsMixed() bool
}

// NewGame - create game
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

// WithDefaults - load default options for game
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

// IsMixed - false
func (game *Game) IsMixed() bool {
	return false
}

// String - game to string
func (game *Game) String() string {
	return fmt.Sprintf("%s %s", game.Type, game.Limit)
}

// Proto - game to proto
func (game *Game) Proto() *message.Game {
	return &message.Game{
		Type: message.GameType(
			message.GameType_value[string(game.Type)],
		).Enum(),

		Limit: message.GameLimit(
			message.GameLimit_value[string(game.Limit)],
		).Enum(),
	}
}

// NewMix - create mix
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

// WithDefaults - load default options for mix
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

// String - mix to string
func (mix *Mix) String() string {
	return fmt.Sprintf("%s", mix.Type)
}

// IsMixed - false
func (mix *Mix) IsMixed() bool {
	return true
}
