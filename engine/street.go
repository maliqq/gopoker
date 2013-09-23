package engine

import (
	"gopoker/engine/street"
)

type Street struct {
	street.Type
	Stages
}
