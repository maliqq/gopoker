package message

import (
	"code.google.com/p/goprotobuf/proto"
)

func NewPlayStart(play *Play) *Message {
	return NewMessage(PlayStart{
		Play: play,
	})
}

func NewStreetStart(name string) *Message {
	return NewMessage(StreetStart{
		Name: proto.String(name),
	})
}

/*
func NewChangeGame(g *model.Game) *Message {
	return NewMessage(ChangeGame{
		Type:  GameType(GameType_value[string(g.Type)]).Enum(),
		Limit: GameLimit(GameLimit_value[string(g.Limit)]).Enum(),
	})
}
*/
