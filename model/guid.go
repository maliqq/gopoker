package model

import (
	"gopoker/util"
)

type Guid string

func (guid Guid) String() string {
	return string(guid)
}

func RandomGuid() Guid {
	return Guid(util.RandomUuid())
}
