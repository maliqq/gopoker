package model

import (
	"gopoker/model/seat"
	"testing"
)

func TestTableMoveButton(t *testing.T) {
	table := NewTable(9)
	t.Logf("button=%d", table.Button)
	table.Button = 8
	t.Logf("button=%d", table.Button)
	table.MoveButton()
	t.Logf("button=%d", table.Button)
	if table.Button != 0 {
		t.FailNow()
	}
}

func TestTableAddPlayer(t *testing.T) {
	table := NewTable(2)

	a := NewPlayer("A")
	b := NewPlayer("B")
	c := NewPlayer("C")

	s, err := table.AddPlayer(a, 0, 10.)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if s.Stack != 10. {
		t.Fatalf("seat.Stack = %.2f", s.Stack)
	}

	s, err = table.AddPlayer(b, 0, 10.)

	t.Logf("error: %v", err)
	if err == nil {
		t.FailNow()
	}

	s, err = table.AddPlayer(b, 1, 12.)

	if s.Stack != 12. {
		t.Fatalf("seat.Stack = %.2f", s.Stack)
	}

	s, err = table.RemovePlayer(a)

	if s.State != seat.Empty {
		t.Fatalf("seat.State = %s", s.State)
	}

	s, err = table.AddPlayer(c, 0, 12.)

	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestTablePosition(t *testing.T) {
	table := NewTable(2)

	emptySeats := table.Seats.From(0).Where(func(s *Seat) bool {
		return s.State == seat.Empty
	})

	if len(emptySeats) != 2 {
		t.FailNow()
	}
}
