package context

import (
	"gopoker/model"
)

// GameRotation - game rotation context
type GameRotation struct {
	*model.Mix
	every   int
	index   int
	counter int
}

const (
	// RotateEvery - number of games to play for switch
	RotateEvery = 8
)

// NewGameRotation - create rotation context
func NewGameRotation(mix *model.Mix, every int) *GameRotation {
	if every == 0 {
		every = RotateEvery
	}
	return &GameRotation{
		Mix:   mix,
		every: every,
	}
}

// Next - get next rotated game
func (rotation *GameRotation) Next() *model.Game {
	if rotation.counter >= rotation.every {
		rotation.counter = 0
		rotation.rotate()
		return rotation.Current()
	}

	rotation.counter++
	return nil
}

func (rotation *GameRotation) rotate() {
	rotation.index++
	if rotation.index >= len(rotation.Mix.Games) {
		rotation.index = 0
	}
}

// Current - get current game in rotation
func (rotation *GameRotation) Current() *model.Game {
	return rotation.Mix.Games[rotation.index]
}
