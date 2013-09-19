package protobuf

import (
	"code.google.com/p/goprotobuf/proto"
)

// NotifyPlayStart - notify new play start
func NewPlayStart(game *Game, stake *Stake, table *Table) *Message {
	return &Message{
		PlayStart: &PlayStart{
			Game:  game,
			Stake: stake,
			Table: table,
		},
	}
}

// NotifyStreetStart - notify new street start
func NewStreetStart(name string) *Message {
	return &Message{
		StreetStart: &StreetStart{
			Name: proto.String(name),
		},
	}
}

// NotifyPlayStop - notify play stop
func NewPlayStop() *Message {
	return &Message{
		PlayStop: &PlayStop{},
	}
}

/*
func NewChangeGame(g *model.Game) *Message {
	return NewMessage(ChangeGame{
		Type:  GameType(GameType_value[string(g.Type)]).Enum(),
		Limit: GameLimit(GameLimit_value[string(g.Limit)]).Enum(),
	})
}
*/
