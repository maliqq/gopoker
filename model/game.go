package model

import (
	"encoding/json"
	"fmt"
	"log"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/exch/message"
	"gopoker/model/game"
	"gopoker/poker/hand"
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
	TableSize    int
	*GameOptions `json:"-"`
}

// Variation - union of Game and Mix
type Variation interface {
	IsMixed() bool
}

// NewGame - create game
func NewGame(gameType game.Type, limit game.Limit, tableSize int) *Game {
	limitedGame, success := gameType.(game.LimitedGame)

	if !success {
		log.Printf("got: %#v", gameType)
		panic("can't create game")
	}

	options := gameOptions(limitedGame)

	maxTableSize := options.MaxTableSize
	if tableSize == 0 || tableSize > maxTableSize {
		tableSize = maxTableSize
	}

	if limit == "" {
		limit = options.DefaultLimit
	}

	game := &Game{
		Type:        limitedGame,
		Limit:       limit,
		TableSize:   tableSize,
		GameOptions: options,
	}

	return game
}

func gameOptions(limitedGame game.LimitedGame) *GameOptions {
	gameOptions, success := Games[limitedGame]
	if !success {
		log.Printf("got: %#v", limitedGame)
		panic("can't find options")
	}

	return gameOptions
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

		TableSize: proto.Int32(int32(game.TableSize)),
	}
}

func (g *Game) UnmarshalProto(protoGame *message.Game) {
	newGame := NewGame(
		game.LimitedGame(protoGame.GetType().String()),
		game.Limit(protoGame.GetLimit().String()),
		int(protoGame.GetTableSize()),
	)
	*g = *newGame
}

func (g *Game) UnmarshalJSON(data []byte) error {
	var result struct {
		Type      string
		Limit     string
		TableSize int
	}

	err := json.Unmarshal(data, &result)

	if err != nil {
		return err
	}

	*g = *NewGame(game.LimitedGame(result.Type), game.Limit(result.Limit), result.TableSize)

	return nil
}
