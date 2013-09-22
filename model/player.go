package model

// Player - player id
type Player Guid

// String - player id
func (player Player) String() string {
	return string(player)
}
