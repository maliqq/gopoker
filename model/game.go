package model

import (
	"encoding/json"
	"os"
	"path"
	"log"
	"fmt"
	"flag"
)

import (
	"gopoker/model/game"
	"gopoker/poker/ranking"
)

const (
	MixedGameMaxTableSize = 8
	GamesConfigFile = "games.json"
	MixesConfigFile = "mixes.json"
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

	Hi ranking.Ranking
	Lo ranking.Ranking

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

var Games map[game.LimitedGame]*GameOptions
var Mixes map[game.MixedGame][]*MixOptions

func readConfig(filepath string, result interface{}) {
	f, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Can't open %s: %s", filepath, err)
	}

	decoder := json.NewDecoder(f)
	
	err = decoder.Decode(&result)
	if err != nil {
		log.Fatal("Can't decode %s: %s", filepath, err)
	}
}

var configDir = flag.String("config-dir", "/etc/gopoker", "Config dir")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	readConfig(path.Join(*configDir, GamesConfigFile), &Games)
	readConfig(path.Join(*configDir, MixesConfigFile), &Mixes)
}

func NewGame(g game.Type, limit game.Limit) *Game {
	limitedGame, success := g.(game.LimitedGame)

	if !success {
		fmt.Printf("got: %s\n", g)
		panic("can't create game")
	}

	game := &Game{
		Type:  limitedGame,
		Limit: limit,
	}

	return game.WithDefaults()
}

func (game *Game) WithDefaults() *Game {
	game.GameOptions = Games[game.Type]

	if game.Limit == "" {
		game.Limit = game.GameOptions.DefaultLimit
	}

	return game
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

	mix := &Mix{
		Type: mixedGame,
	}

	return mix.WithDefaults()
}

func (mix *Mix) WithDefaults() *Mix {
	options, _ := Mixes[mix.Type]

	games := make([]*Game, len(options))
	for i, mixOptions := range options {
		games[i] = NewGame(mixOptions.Type, mixOptions.Limit)
	}
	mix.Games = games

	return mix
}

func (mix *Mix) String() string {
	return fmt.Sprintf("%s %s", mix.Type)
}

func (mix *Mix) IsMixed() bool {
	return true
}
