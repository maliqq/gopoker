package model

import (
	"code.google.com/p/goprotobuf/proto"
)

// Player - player id
type Player string

// RouteKey - route key for broker
func (player Player) RouteKey() string {
	return string(player)
}

// Proto - player to protobuf
func (player Player) Proto() *string {
	return proto.String(string(player))
}
