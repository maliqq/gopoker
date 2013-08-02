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

type seatSlice struct {
	from  int
	seats *Seats
}

type seatFilter func(s *Seat) bool

// From - seats from position
func (seats Seats) From(from int) seatSlice {
	return seatSlice{from, &seats}
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

func (slice seatSlice) len() int {
	return len(*slice.seats)
}

// All - get all seats
func (slice seatSlice) All() []int {
	index := []int{}
	for i := slice.from + 1; i < slice.len(); i++ {
		index = append(index, i)
	}
	for i := 0; i <= slice.from; i++ {
		index = append(index, i)
	}
	return index
}

// Where - get seats matching filter
func (slice seatSlice) Where(filter seatFilter) []int {
	result := []int{}

	for _, pos := range slice.All() {
		seat := (*slice.seats)[pos]
		if filter(seat) {
			result = append(result, pos)
		}
	}

	return result
}

// Active - get all active seats
func (slice seatSlice) Active() []int {
	return slice.Where(func(s *Seat) bool {
		return s.State == seat.Play || s.State == seat.PostBB
	})
}

// Waiting - get all waiting seats
func (slice seatSlice) Waiting() []int {
	return slice.Where(func(s *Seat) bool {
		return s.State == seat.WaitBB
	})
}

// Playing - get all playing seats
func (slice seatSlice) Playing() []int {
	return slice.Where(func(s *Seat) bool {
		return s.State == seat.Play
	})
}

// InPlay - get all seats in play
func (slice seatSlice) InPlay() []int {
	return slice.Where(func(s *Seat) bool {
		return s.State == seat.Play || s.State == seat.Bet
	})
}

// InPot - get all seats in pot
func (slice seatSlice) InPot() []int {
	return slice.Where(func(s *Seat) bool {
		return s.State == seat.Play || s.State == seat.Bet || s.State == seat.AllIn
	})
}
