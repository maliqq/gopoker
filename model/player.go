package model

import (
	"code.google.com/p/goprotobuf/proto"
)

type Player string

func (player Player) RouteKey() string {
	return string(player)
}

func (player Player) Proto() *string {
	return proto.String(string(player))
}
