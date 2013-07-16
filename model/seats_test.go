package model

import (
	"testing"
)

func TestSeats(t *testing.T) {
	seats := Seats{&Seat{Stack: 1}, &Seat{Stack: 2}, &Seat{Stack: 3}, &Seat{Stack: 4}, &Seat{Stack: 5}, &Seat{Stack: 6}}

	for i := 0; i < len(seats); i++ {
		t.Logf("seats from %d=%#v", i, seats.From(i).All())
	}
}
