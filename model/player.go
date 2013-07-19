package model

type Player string

func (player Player) RouteKey() string {
	return string(player)
}
