package model

import (
	"testing"
)

func TestPot(t *testing.T) {
	pot := NewPot()

	pot.Add("1", 1., false)
	pot.Add("2", 2., false)
	pot.Add("3", 4.25, false)
	pot.Add("1", 3.25, false)
	pot.Add("2", 2.25, false)

	t.Logf("%s", pot)

	if pot.Total() != 4.25*3 {
		t.FailNow()
	}

	pot1 := NewPot()

	pot1.Add("1", 3., false)
	pot1.Add("2", 3., false)
	pot1.Add("3", 1., true)

	t.Logf("%s", pot1)

	if len(pot1.Side) != 1 {
		t.FailNow()
	}

	pot2 := NewPot()

	pot2.Add("1", 2., false)
	pot2.Add("2", 4., true)
	pot2.Add("3", 4., false)
	pot2.Add("1", 2., false)

	t.Logf("%s", pot2)

	if len(pot2.Side) != 1 {
		t.FailNow()
	}

	pot3 := NewPot()

	pot3.Add("1", 1., false)
	pot3.Add("2", 4., true)
	pot3.Add("3", 3., true)
	pot3.Add("1", 3., false)

	t.Logf("%s", pot3)

	if len(pot3.Side) != 2 {
		t.FailNow()
	}
}
