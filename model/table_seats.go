package model

import (
	"fmt"
)

import (
	"gopoker/model/seat"
)

// Seats - list of Seat
type Seats []*Seat

// NewSeats - create seats of specified size
func NewSeats(size int) []*Seat {
	seats := make([]*Seat, size)
	for i := 0; i < size; i++ {
		seats[i] = NewSeat()
	}

	return seats
}

func (seats Seats) At(pos int) *Seat {
	return seats[pos]
}

// String - seats to string
func (seats Seats) String() string {
	str := ""

	for i, tableSeat := range seats {
		str += fmt.Sprintf("seat %d: ", i+1)
		if tableSeat.State == seat.Empty {
			str += "empty\n"
		} else {
			str += fmt.Sprintf("%s (%.2f/%.2f) %s\n", tableSeat.Player, tableSeat.Stack, tableSeat.Bet, tableSeat.State)
		}
	}

	return str
}
