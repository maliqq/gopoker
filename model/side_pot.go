package model

import (
	"fmt"
)

// SidePot - side pot
type SidePot struct {
	Barrier float64
	Members map[Player]float64
}

// NewSidePot - create side pot with barrier amount
func NewSidePot(amount float64) *SidePot {
	return &SidePot{
		Barrier: amount,
		Members: map[Player]float64{},
	}
}

// String - side pot to string
func (pot *SidePot) String() string {
	s := "SIDE POT\n"
	for member, value := range pot.Members {
		s += fmt.Sprintf("#%s: %.2f\n", member, value)
	}
	s += fmt.Sprintf("-- Total: %.2f Barrier: %.2f\n", pot.Total(), pot.Barrier)
	return s
}

// Total - get side pot total
func (pot *SidePot) Total() float64 {
	sum := 0.
	for _, amount := range pot.Members {
		sum += amount
	}

	return sum
}

// IsActive - check if side pot is active
func (pot *SidePot) IsActive() bool {
	return len(pot.Members) > 0 && pot.Total() > 0.
}

// Add - add amount to side pot
func (pot *SidePot) Add(member Player, amount float64) float64 {
	if amount > 0. {
		value, exists := pot.Members[member]
		if !exists {
			value = 0.
			pot.Members[member] = value
		}

		barrier := pot.Barrier
		if barrier == 0. {
			pot.Members[member] += amount
			return 0.
		}

		if barrier >= amount {
			pot.Members[member] = barrier
			return value + amount - barrier
		}
	}

	return amount
}

// Split - split side pot by barrier
func (pot *SidePot) Split(member Player, remain float64) (*SidePot, *SidePot) {
	pot.Members[member] += remain

	bet := pot.Members[member]

	main := NewSidePot(0.)
	side := NewSidePot(bet)

	for key, value := range pot.Members {
		if value > bet {
			if key != member {
				main.Members[key] = value - bet
			}
			side.Members[key] = bet
		} else {
			side.Members[key] = value
		}
	}

	return main, side
}
