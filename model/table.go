package model

import (
	_ "container/ring"
	"fmt"
	_ "net"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model/position"
	"gopoker/protocol/message"
)

type Table struct {
	Size   int // number of seats
	Button int

	Seats   Seats
	seating map[Player]int
}

func NewTable(size int) *Table {
	return &Table{
		Size:    size,
		Seats:   NewSeats(size),
		seating: map[Player]int{},
	}
}

func (t *Table) MoveButton() {
	t.SetButton(t.Button + 1)
}

func (t *Table) SetButton(pos int) {
	t.Button = position.Cycle(pos, t.Size)
}

func (t *Table) Register(player Player, pos int) {
	t.seating[player] = pos
}

func (t *Table) Unregister(player Player) {
	delete(t.seating, player)
}

func (t *Table) Seating(player Player) (int, bool) {
	pos, found := t.seating[player]

	return pos, found
}

func (t *Table) AddSeating(player Player, pos int) (*Seat, error) {
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

func (t *Table) Pos(player Player) (int, error) {
	pos, found := t.seating[player]
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

func (t *Table) AllPlayers() []Player {
	players := []Player{}
	for player, _ := range t.seating {
		players = append(players, player)
	}
	return players
}

func (t *Table) Player(pos int) Player {
	return t.Seat(pos).Player
}

func (t *Table) AddPlayer(player Player, pos int, amount float64) (*Seat, error) {
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

func (t *Table) RemovePlayer(player Player) (*Seat, error) {
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

func (t *Table) Proto() *message.Table {
	seats := make([]*message.Seat, t.Size)
	for i, seat := range t.Seats {
		seats[i] = seat.Proto()
	}
	return &message.Table{
		Size:   proto.Int32(int32(t.Size)),
		Button: proto.Int32(int32(t.Button)),
		Seats:  seats,
	}
}
