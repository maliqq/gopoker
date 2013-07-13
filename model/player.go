package model

type Player struct {
	Id Id
}

func NewPlayer(id Id) *Player {
	return &Player{Id: id}
}

func (player *Player) String() string {
	return string(player.Id)
}

func (player *Player) RouteKey() string {
	return string(player.Id)
}
