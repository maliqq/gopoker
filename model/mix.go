package model

import (
	"fmt"
	"log"
)

import (
	"gopoker/model/game"
)

const (
	// MixedGameMaxTableSize - table size for mixed game
	MixedGameMaxTableSize = 8
)

// MixOptions - mix options
type MixOptions struct {
	Type game.LimitedGame
	game.Limit
}

// Mix - mix
type Mix struct {
	Type  game.MixedGame
	Games []*Game `json:"-"`
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
