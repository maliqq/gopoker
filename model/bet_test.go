package model

import (
	"testing"
)

func TestBet(t *testing.T) {
	b := NewCheck()
	if b.PrintString() != "check" {
		t.Fatalf("bet=%#v", b)
	}

	b = NewRaise(10.0)
	if b.PrintString() != "raise 10.00" {
		t.Fatalf("bet=%#v (%s)", b, b)
	}

	b = NewCall(10.0)
	if b.PrintString() != "call 10.00" {
		t.Fatalf("bet=%#v (%s)", b, b)
	}

	b = NewFold()
	if b.PrintString() != "fold" {
		t.Fatalf("bet=%#v (%s)", b, b)
	}
}
