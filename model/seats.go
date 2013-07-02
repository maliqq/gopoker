package model

import (
	"gopoker/model/seat"
)

type Seats []*Seat

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

func (seats Seats) From(from int) seatSlice {
	return seatSlice{from, &seats}
}

func (slice seatSlice) Where(filter seatFilter) []int {
	seats := []int{}

	for pos, seat := range *slice.seats {
		if filter(seat) {
			seats = append(seats, pos)
		}
	}

	return seats
}

func (slice seatSlice) Active() []int {
	return slice.Where(func(s *Seat) bool {
		return s.State == seat.Play || s.State == seat.PostBB
	})
}

func (slice seatSlice) Waiting() []int {
	return slice.Where(func(s *Seat) bool {
		return s.State == seat.WaitBB
	})
}

func (slice seatSlice) InPlay() []int {
	return slice.Where(func(s *Seat) bool {
		return s.State == seat.Play
	})
}

func (slice seatSlice) InPot() []int {
	return slice.Where(func(s *Seat) bool {
		return s.State == seat.Play || s.State == seat.AllIn || s.State == seat.Bet
	})
}
