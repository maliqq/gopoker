package context

import (
	"gopoker/model"
	"gopoker/model/game"
	"testing"
)

func TestBetting(t *testing.T) {
	b := NewBetting()

	player := model.Player("A")
	seat := &model.Seat{
		Player: player,
		Stack:  200.,
	}
	stake := model.NewStake(10.)
	t.Logf("Stake: %s", stake)

	// no limit
	g := model.NewGame(game.Texas, game.NoLimit)

	b.RequireBet(0, seat, g.Limit, stake)
	req := b.BetRange

	t.Logf("Required: %s", req)

	if req.Call != 0. {
		t.Fatalf("required.Call mismatch")
	}

	if req.Min != stake.BigBlindAmount() {
		t.Fatalf("required.Min mismatch")
	}

	if req.Max != seat.Stack {
		t.Fatalf("required.Max mismatch")
	}

	// fixed limit
	g = model.NewGame(game.Texas, game.FixedLimit)

	b.RequireBet(0, seat, g.Limit, stake)
	req = b.BetRange

	t.Logf("Required: %s", req)

	if req.Call != 0. {
		t.Fatalf("required.Call mismatch")
	}

	if req.Min != stake.BigBlindAmount() {
		t.Fatalf("required.Min mismatch")
	}

	if req.Max != stake.BigBlindAmount() {
		t.Fatalf("required.Max mismatch")
	}

	// pot limit
	g = model.NewGame(game.Omaha, game.PotLimit)

	b.RequireBet(0, seat, g.Limit, stake)
	req = b.BetRange

	t.Logf("Required: %s", req)

	if req.Call != 0. {
		t.Fatalf("required.Call mismatch")
	}

	if req.Min != stake.BigBlindAmount() {
		t.Fatalf("required.Min mismatch")
	}

	if req.Max != 0. {
		t.Fatalf("required.Max mismatch")
	}
}
