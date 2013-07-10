package context

import (
	"gopoker/model"
)

type GameRotation struct {
	*model.Mix
	every   int
	index   int
	counter int
}

const (
	RotateEvery = 8
)

func NewGameRotation(mix *model.Mix, every int) *GameRotation {
	if every == 0 {
		every = RotateEvery
	}
	return &GameRotation{
		Mix:   mix,
		every: every,
	}
}

func (rotation *GameRotation) Next() *model.Game {
	if rotation.counter >= rotation.every {
		rotation.counter = 0
		rotation.Rotate()
	} else {
		rotation.counter++
	}
	return rotation.Current()
}

func (rotation *GameRotation) Rotate() {
	rotation.index++
	if rotation.index >= len(rotation.Mix.Games) {
		rotation.index = 0
	}
}

func (rotation *GameRotation) Current() *model.Game {
	return rotation.Mix.Games[rotation.index]
}
