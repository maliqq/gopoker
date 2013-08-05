package model

import (
	"testing"
)

func TestBet(t *testing.T) {
	b := NewCheck()
	if b.PrintString() != "Check" {
		t.Fatalf("bet=%#v", b)
	}

	b = NewRaise(10.0)
	if b.PrintString() != "Raise 10.00" {
		t.Fatalf("bet=%#v (%s)", b, b)
	}

	b = NewCall(10.0)
	if b.PrintString() != "Call 10.00" {
		t.Fatalf("bet=%#v (%s)", b, b)
	}

	b = NewFold()
	if b.PrintString() != "Fold" {
		t.Fatalf("bet=%#v (%s)", b, b)
	}
}
