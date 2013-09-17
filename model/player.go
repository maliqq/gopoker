package model

import (
	"code.google.com/p/goprotobuf/proto"
)

// Player - player id
type Player Guid

// String - player id
func (player Player) String() string {
	return string(player)
}

// Proto - player to protobuf
func (player Player) Proto() *string {
	return proto.String(string(player))
}
