package model

import (
	"fmt"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/model/bet"
	seatState "gopoker/model/seat"
	"gopoker/protocol/message"
)

type Seat struct {
	Player Player

	State seatState.State
	Stack float64

	Bet float64
}

func NewSeat() *Seat {
	return &Seat{State: seatState.Empty}
}

func (seat *Seat) String() string {
	return fmt.Sprintf("state=%s player=%s stack=%.2f bet=%.2f",
		seat.State,
		seat.Player,
		seat.Stack,
		seat.Bet,
	)
}

func (seat *Seat) Clear() {
	seat.State = seatState.Empty
	seat.Player = ""
	seat.Stack = 0.
	seat.Bet = 0.
}

func (seat *Seat) Play() {
	seat.Bet = 0.
	seat.State = seatState.Play
}

func (seat *Seat) Check() {
	seat.State = seatState.Bet
}

func (seat *Seat) Fold() {
	seat.Bet = 0.
	seat.State = seatState.Fold
}

func (seat *Seat) Calls(amount float64) bool {
	return seat.Bet >= amount || seat.State == seatState.AllIn
}

func (seat *Seat) SetBet(amount float64) {
	seat.Stack += (seat.Bet - amount)
	seat.Bet = amount

	if seat.Stack == 0. {
		seat.State = seatState.AllIn
	} else {
		seat.State = seatState.Bet
	}
}

func (seat *Seat) ForceBet(amount float64) {
	seat.Stack -= amount
	seat.Bet = amount

	if seat.Stack == 0. {
		seat.State = seatState.AllIn
	} else {
		seat.State = seatState.Play
	}
}

func (seat *Seat) AddBet(b *Bet) (float64, bool) {
	var put float64

	switch b.Type {
	case bet.Fold:
		put = 0
		seat.Fold()

	case bet.Check:
		put = 0
		seat.Check()

	case bet.Call, bet.Raise:
		put = b.Amount - seat.Bet
		seat.SetBet(b.Amount)

	default:
		put = b.Amount - seat.Bet
		if b.IsForced() {
			seat.ForceBet(b.Amount)
		}
	}

	return put, seat.State == seatState.AllIn
}

func (seat *Seat) SetPlayer(player Player) error {
	if seat.State != seatState.Empty {
		return fmt.Errorf("Seat is not empty.")
	}

	seat.Player = player
	seat.State = seatState.Taken

	return nil
}

func (seat *Seat) SetStack(amount float64) {
	seat.Stack = amount

	if seat.State == seatState.Taken {
		seat.State = seatState.Ready
	}
}

func (seat *Seat) AdvanceStack(amount float64) {
	seat.Stack += amount
}

func (seat *Seat) Proto() *message.Seat {
	return &message.Seat{
		State: message.SeatState(
			message.SeatState_value[string(seat.State)],
		).Enum(),

		Stack: proto.Float64(seat.Stack),

		Bet: proto.Float64(seat.Bet),
	}
}
