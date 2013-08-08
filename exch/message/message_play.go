package message

import (
	"code.google.com/p/goprotobuf/proto"
)

// NotifyPlayStart - notify new play start
func NotifyPlayStart(play *Play) *Message {
	return NewMessage(PlayStart{
		Play: play,
	})
}

// NotifyStreetStart - notify new street start
func NotifyStreetStart(name string) *Message {
	return NewMessage(StreetStart{
		Name: proto.String(name),
	})
}

// NotifyPlayStop - notify play stop
func NotifyPlayStop() *Message {
	return NewMessage(PlayStop{})
}

/*
func NotifyChangeGame(g *model.Game) *Message {
	return NewMessage(ChangeGame{
		Type:  GameType(GameType_value[string(g.Type)]).Enum(),
		Limit: GameLimit(GameLimit_value[string(g.Limit)]).Enum(),
	})
}
*/
