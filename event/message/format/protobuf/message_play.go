package protobuf

import (
	"code.google.com/p/goprotobuf/proto"
)

// NotifyPlayStart - notify new play start
func NewPlayStart(play *Play) *Message {
	return &Message{
		Payload: &Payload{
			PlayStart: &PlayStart{
				Play: play,
			},
		},
	}
}

// NotifyStreetStart - notify new street start
func NewStreetStart(name string) *Message {
	return &Message{
		Payload: &Payload{
			StreetStart: &StreetStart{
				Name: proto.String(name),
			},
		},
	}
}

// NotifyPlayStop - notify play stop
func NewPlayStop() *Message {
	return &Message{
		Payload: &Payload{
			PlayStop: &PlayStop{},
		},
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
