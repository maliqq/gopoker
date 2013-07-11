package model

import (
	"gopoker/util"
)

type Id string

func RandomId() Id {
	return Id(util.RandomUuid())
}
