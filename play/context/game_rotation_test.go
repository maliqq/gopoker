package context

import (
	"testing"
)

import (
	"gopoker/model"
	"gopoker/model/game"
)

func TestGameRotation(t *testing.T) {
	stake := game.NewStake(10.)
	mix := model.NewMixedGame(game.Horse, stake)
	t.Logf("mix=%s", mix)

	rotation := NewGameRotation(mix, 0)

	var g *model.Game
	for i := 0; i < RotateEvery; i++ {
		g = rotation.Next()
		t.Logf("game=%s", g)
	}
	if g.Type != mix.Games[0].Type {
		t.Fatalf("should be still first game")
	}
	g = rotation.Next()
	t.Logf("game after N=%d: %s", RotateEvery, g)
	if g.Type != mix.Games[1].Type {
		t.Fatalf("should switch to next game")
	}
}
