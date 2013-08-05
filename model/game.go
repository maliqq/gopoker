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

// Game - game
type Game struct {
	Type game.LimitedGame
	game.Limit
	*GameOptions `json:"-"`
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
