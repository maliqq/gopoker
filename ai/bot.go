package ai

import (
	"gopoker/client/zeromq_client"
	"gopoker/util"
)

type Bot struct {
	Id        string
	zmqConn *zeromq_client.Connection
}

func NewBot(sockAddr string) *Bot {
	id := util.RandomUuid()

	return &Bot{
		Id:        id,
		zmqConn: zeromq_client.NewConnection(sockAddr),
	}
}
