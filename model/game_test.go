package model

import (
	"encoding/json"
	"gopoker/model/game"
	"reflect"
	"testing"
)

func TestGameMarshal(t *testing.T) {
	original := NewGame(game.Texas, game.NoLimit, 9)

	s, _ := json.Marshal(original)

	var game *Game
	err := json.Unmarshal(s, &game)
	if err != nil {
		t.Fatal("unmarshal error: ", err)
	}

	if !reflect.DeepEqual(game.GameOptions, original.GameOptions) {
		t.Fatalf("game options mismatch: %#v vs %#v", game.GameOptions, original.GameOptions)
	}
}

func TestMixMarshal(t *testing.T) {
	original := NewMix(game.Horse, 8)

	s, _ := json.Marshal(original)

	var mix *Mix
	err := json.Unmarshal(s, &mix)
	if err != nil {
		t.Fatal("unmarshal error: ", err)
	}

	if !reflect.DeepEqual(mix.Games, original.Games) {
		t.Fatalf("games mismatch: %#v vs %#v", mix.Games, original.Games)
	}
}
