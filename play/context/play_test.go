package context

import (
	"encoding/json"
	"testing"
)

import (
	"gopoker/model"
	"gopoker/model/game"
)

func TestPlayMarshalJSON(t *testing.T) {
	stake := game.NewStake(10.)
	table := model.NewTable(9)

	play := NewPlay(model.NewGame(game.Texas, game.NoLimit, stake), table)

	result, err := json.Marshal(play)

	if err != nil {
		t.Fatalf("error marshaling: %s", err)
	}

	t.Logf("%s", result)
	t.FailNow()
}
