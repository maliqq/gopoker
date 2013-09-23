package model

type Seating map[Player]int

func NewSeating() Seating {
	return Seating{}
}

func (seating Seating) Add(p Player, pos int) {
	seating[p] = pos
}

func (seating Seating) Remove(p Player) {
	delete(seating, p)
}

func (seating Seating) Check(player Player) (int, bool) {
	pos, presence := seating[player]
	return pos, presence
}
