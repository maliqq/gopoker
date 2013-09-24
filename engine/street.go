package engine

import (
	"gopoker/engine/street"
)

type Street struct {
	street.Type
	StageStrategy
}

func (process Street) String() string {
	return string(process.Type)
}

type StreetStrategy []Street
