package model

import (
	"fmt"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/event/message/format/protobuf"
	"gopoker/model/position"
)

// Table - players seating
type Table struct {
	Size   int // number of seats
	Button int

	Seats   Seats
	seating map[Player]int
}

// NewTable - create table by size
func NewTable(size int) *Table {
	return &Table{
		Size:    size,
		Seats:   NewSeats(size),
		seating: map[Player]int{},
	}
}

// MoveButton - move button to next position
func (t *Table) MoveButton() {
	t.SetButton(t.Button + 1)
}

// SetButton - set button to specified position
func (t *Table) SetButton(pos int) {
	t.Button = position.Cycle(pos, t.Size)
}

// Register - register user at specified position
func (t *Table) Register(player Player, pos int) {
	t.seating[player] = pos
}

// Unregister - unregister user from specified position
func (t *Table) Unregister(player Player) {
	delete(t.seating, player)
}

// Seating - get player seating
func (t *Table) Seating(player Player) (int, bool) {
	pos, found := t.seating[player]

	return pos, found
}

// AddSeating - add player seating at pos
func (t *Table) AddSeating(player Player, pos int) (*Seat, error) {
	seat := t.Seat(pos)

	err := seat.SetPlayer(player)
	if err != nil {
		return nil, err
	}

	t.Register(player, pos)

	return seat, nil
}

// RemoveSeating - remove seating at pos
func (t *Table) RemoveSeating(pos int) (*Seat, error) {
	seat := t.Seat(pos)

	t.Unregister(seat.Player)
	seat.Clear()

	return seat, nil
}

// Pos - get position for player
func (t *Table) Pos(player Player) (int, error) {
	pos, found := t.seating[player]
	if !found {
		return 0, fmt.Errorf("Player not found.")
	}
	return pos, nil
}

// Seat - get seat for position
func (t *Table) Seat(pos int) *Seat {
	return t.Seats[pos]
}

// AllSeats - get all seats for table
func (t *Table) AllSeats() seatSlice {
	return t.Seats.From(t.Button)
}

// AllPlayers - get all players for table
func (t *Table) AllPlayers() []Player {
	players := []Player{}
	for player := range t.seating {
		players = append(players, player)
	}
	return players
}

// Player - get player for position
func (t *Table) Player(pos int) Player {
	return t.Seat(pos).Player
}

// AddPlayer - add player with amount of stack at position
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

// RemovePlayer - remove player from table
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

// String - table to string
func (t *Table) String() string {
	return fmt.Sprintf("size: %d button: %d\n%s", t.Size, t.Button, t.Seats)
}

// Proto - table to protobuf
func (t *Table) Proto() *protobuf.Table {
	seats := make([]*protobuf.Seat, t.Size)
	for i, seat := range t.Seats {
		seats[i] = seat.ProtoSeat()
	}
	return &protobuf.Table{
		Size:   proto.Int32(int32(t.Size)),
		Button: proto.Int32(int32(t.Button)),
		Seats:  seats,
	}
}
