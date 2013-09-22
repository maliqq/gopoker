package engine

import (
	"gopoker/engine/street"
)

type StreetProcess struct {
	Street street.Type
	Stages StageProcesses
}
