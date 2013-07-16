package model

import (
	"fmt"
	_ "net"
)

import (
	"gopoker/model/position"
)

type Table struct {
	Size   int // number of seats
	Button int

	Seats   Seats
	seating map[Id]int

	//Waiting  []*Player
	//Watchers []net.Conn
}

func NewTable(size int) *Table {
	return &Table{
		Size:    size,
		Seats:   NewSeats(size),
		seating: map[Id]int{},
	}
}

func (t *Table) MoveButton() {
	t.SetButton(t.Button + 1)
}

func (t *Table) SetButton(pos int) {
	t.Button = position.Cycle(pos, t.Size)
}

func (t *Table) Register(player *Player, pos int) {
	t.seating[player.Id] = pos
}

func (t *Table) Unregister(player *Player) {
	delete(t.seating, player.Id)
}

func (t *Table) Seating(player *Player) (int, bool) {
	pos, found := t.seating[player.Id]

	return pos, found
}

func (t *Table) AddSeating(player *Player, pos int) (*Seat, error) {
	seat := t.Seat(pos)

	err := seat.SetPlayer(player)
	if err != nil {
		return nil, err
	}

	t.Register(player, pos)

	return seat, nil
}

func (t *Table) RemoveSeating(pos int) (*Seat, error) {
	seat := t.Seat(pos)

	t.Unregister(seat.Player)
	seat.Clear()

	return seat, nil
}

func (t *Table) Pos(player *Player) (int, error) {
	pos, found := t.seating[player.Id]
	if !found {
		return 0, fmt.Errorf("Player not found.")
	}
	return pos, nil
}

func (t *Table) Seat(pos int) *Seat {
	return t.Seats[pos]
}

func (t *Table) AllSeats() seatSlice {
	return t.Seats.From(t.Button)
}

func (t *Table) Player(pos int) *Player {
	return t.Seat(pos).Player
}

func (t *Table) AddPlayer(player *Player, pos int, amount float64) (*Seat, error) {
	oldPos, hasPlayer := t.Seating(player)

	if hasPlayer {
		return nil, fmt.Errorf("Player already joined at %d", oldPos)
	}

	seat, err := t.AddSeating(player, pos)

	if err != nil {
		return nil, err
	}

	if amount != 0. {
		seat.SetStack(amount)
	}

	return seat, nil
}

func (t *Table) RemovePlayer(player *Player) (*Seat, error) {
	pos, hasPlayer := t.Seating(player)

	if !hasPlayer {
		return nil, fmt.Errorf("Can't remove player")
	}

	seat, err := t.RemoveSeating(pos)

	if err != nil {
		return nil, err
	}

	return seat, nil
}

func (t *Table) String() string {
	return fmt.Sprintf("size: %d button: %d\n%s", t.Size, t.Button, t.Seats)
}
