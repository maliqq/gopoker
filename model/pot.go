package model

import (
	"fmt"
)

type SidePot struct {
	Barrier float64
	Members map[Id]float64
}

type Pot struct {
	Main *SidePot
	Side []*SidePot
}

func NewSidePot(amount float64) *SidePot {
	return &SidePot{
		Barrier: amount,
		Members: map[Id]float64{},
	}
}

func NewPot() *Pot {
	return &Pot{
		Main: NewSidePot(0.),
		Side: []*SidePot{},
	}
}

func (pot *SidePot) Total() float64 {
	sum := 0.
	for _, amount := range pot.Members {
		sum += amount
	}
	return sum
}

func (pot *Pot) Total() float64 {
	sum := 0.
	for _, sidePot := range pot.SidePots() {
		sum += sidePot.Total()
	}
	return sum
}

func (pot *SidePot) IsActive() bool {
	return len(pot.Members) > 0 && pot.Total() > 0.
}

func (pot *SidePot) Add(member Id, amount float64) float64 {
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

func (pot *SidePot) String() string {
	s := "SIDE POT\n"
	for member, value := range pot.Members {
		s += fmt.Sprintf("#%s: %.2f\n", member, value)
	}
	s += fmt.Sprintf("-- Total: %.2f Barrier: %.2f\n", pot.Total(), pot.Barrier)
	return s
}

func (pot *Pot) String() string {
	s := "\n"
	for _, pot := range pot.SidePots() {
		s += pot.String()
	}
	return s
}

func (pot *Pot) SidePots() []*SidePot {
	//return append(pot.Side, pot.Main)
	pots := []*SidePot{}

	if pot.Main.IsActive() {
		pots = append(pots, pot.Main)
	}

	for _, side := range pot.Side {
		if side.IsActive() {
			pots = append(pots, side)
		}
	}

	return pots
}

func (pot *SidePot) Split(member Id, remain float64) (*SidePot, *SidePot) {
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

func (pot *Pot) Split(member Id, amount float64) {
	main, side := pot.Main.Split(member, amount)
	pot.Side = append(pot.Side, side)
	pot.Main = main
}

func (pot *Pot) Add(member Id, amount float64, allin bool) {
	remain := amount
	for _, side := range pot.Side {
		remain = side.Add(member, remain)
	}

	if allin {
		pot.Split(member, remain)
	} else {
		pot.Main.Add(member, remain)
	}
}
