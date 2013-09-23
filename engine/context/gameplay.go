package context

import (
	"gopoker/model"
)

type Gameplay struct {
	Game  *model.Game
	Stake *model.Stake
	Mix   *model.Mix
	Table *model.Table
}
