package model

type Place struct {
	Name        string
	Country     string
	CountryCode string
	City        string
	Region      string
}

type Player struct {
	Id       Id
	Name     string
	NickName string
	Place    *Place
	Avatar   string
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
