package model

import (
	"fmt"
)

import (
	"code.google.com/p/goprotobuf/proto"
)

import (
	"gopoker/event/message/format/protobuf"
	"gopoker/model/bet"
	seatState "gopoker/model/seat"
)

// Seat - table seat
type Seat struct {
	Player Player

	State seatState.State
	Stack float64

	Bet float64
}

// NewSeat - create seat
func NewSeat() *Seat {
	return &Seat{State: seatState.Empty}
}

// String - seat to string
func (seat *Seat) String() string {
	return fmt.Sprintf("state=%s player=%s stack=%.2f bet=%.2f",
		seat.State,
		seat.Player,
		seat.Stack,
		seat.Bet,
	)
}

// Clear - reset seat
func (seat *Seat) Clear() {
	seat.State = seatState.Empty
	seat.Player = ""
	seat.Stack = 0.
	seat.Bet = 0.
}

// Play - mark seat as playing
func (seat *Seat) Play() {
	seat.Bet = 0.
	seat.State = seatState.Play
}

// Check - seat checked
func (seat *Seat) Check() {
	seat.State = seatState.Bet
}

// Fold - seat folded
func (seat *Seat) Fold() {
	seat.Bet = 0.
	seat.State = seatState.Fold
}

// Calls - check if seat called required amount
func (seat *Seat) Calls(amount float64) bool {
	return seat.Bet >= amount || seat.State == seatState.AllIn
}

// SetBet - seat wager
func (seat *Seat) SetBet(amount float64) {
	seat.Stack += (seat.Bet - amount)
	seat.Bet = amount

	if seat.Stack == 0. {
		seat.State = seatState.AllIn
	} else {
		seat.State = seatState.Bet
	}
}

// ForceBet - force seat to bet amount
func (seat *Seat) ForceBet(amount float64) {
	seat.Stack -= amount
	seat.Bet = amount

	if seat.Stack == 0. {
		seat.State = seatState.AllIn
	} else {
		seat.State = seatState.Play
	}
}

// AddBet - add bet to seat
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

// SetPlayer - assign player to seat
func (seat *Seat) SetPlayer(player Player) error {
	if seat.State != seatState.Empty {
		return fmt.Errorf("Seat is not empty.")
	}

	seat.Player = player
	seat.State = seatState.Taken

	return nil
}

// SetStack - assign stack to seat
func (seat *Seat) SetStack(amount float64) {
	seat.Stack = amount

	if seat.State == seatState.Taken {
		seat.State = seatState.Ready
	}
}

// AdvanceStack - add stack to seat
func (seat *Seat) AdvanceStack(amount float64) {
	seat.Stack += amount
}

// Proto - seat to protobuf
func (seat *Seat) Proto() *protobuf.Seat {
	return &protobuf.Seat{
		State: protobuf.SeatState(
			protobuf.SeatState_value[string(seat.State)],
		).Enum(),

		Stack: proto.Float64(seat.Stack),

		Bet: proto.Float64(seat.Bet),
	}
}

func (seat *Seat) UnmarshalProto(protoSeat *protobuf.Seat) {
	// FIXME
	*seat = Seat{}
}
