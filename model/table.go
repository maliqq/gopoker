package model

import (
	"fmt"
  "container/ring"
)

// Table - players seating
type Table struct {
	Size   int // number of seats
	Button int // button position
	Cursor int // acting player

	Seats   Seats

	ring *ring.Ring
	seating Seating
}

// NewTable - create table by size
func NewTable(size int) *Table {
	seats := NewSeats(size)
	t := Table{
		Size:    size,
		Seats:   seats,
	}
	t.createSeating()
	t.createRing()

	return &t
}

func (t *Table) createRing() {
  r := ring.New(t.Size)
  
  for i := 0; i < t.Size; i++ {
    r.Value = Box{Pos: i, Seat: t.Seats[i]}
    r = r.Next()
  }

  r = r.Move(t.Button)

  t.ring = r
}

func (t *Table) Ring() *Ring {
	return &Ring{t.ring}
}

// MoveButton - move button to next position
func (t *Table) MoveButton() {
	t.ring = t.ring.Move(1)
	t.Button = t.ring.Value.(Box).Pos
}

// SetButton - set button to specified position
func (t *Table) SetButton(pos int) {
	if t.Button != pos {
		t.ring = t.ring.Move(pos - t.Button)
		t.Button = t.ring.Value.(Box).Pos
	}
}

func (t *Table) createSeating() {
	t.seating = NewSeating()
}

func (t *Table) Seating(player Player) (int, bool) {
	return t.seating.Check(player)
}

// AddSeating - add player seating at pos
func (t *Table) AddSeating(player Player, pos int) (*Seat, error) {
	seat := t.Seat(pos)

	err := seat.SetPlayer(player)
	if err != nil {
		return nil, err
	}

	t.seating.Add(player, pos)

	return seat, nil
}

// RemoveSeating - remove seating at pos
func (t *Table) RemoveSeating(pos int) (*Seat, error) {
	seat := t.Seat(pos)

	seat.Clear()
	t.seating.Remove(seat.Player)

	return seat, nil
}

// Pos - get position for player
func (t *Table) Pos(player Player) (int, error) {
	pos, found := t.seating.Check(player)
	if !found {
		return 0, fmt.Errorf("Player not found.")
	}
	return pos, nil
}

// Seat - get seat for position
func (t *Table) Seat(pos int) *Seat {
	return t.Seats.At(pos)
}

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
	return fmt.Sprintf("size: %d button: %d\n%s", t.Size, t.Button + 1, t.Seats)
}
