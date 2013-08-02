package ai

import (
	"gopoker/client/zeromq_client"
	"gopoker/util"
)

// Bot - bot
type Bot struct {
	ID      string
	zmqConn *zeromq_client.Connection
}

// NewBot - create new bot
func NewBot(sockAddr string) *Bot {
	id := util.RandomUuid()

	return &Bot{
		ID:      id,
		zmqConn: zeromq_client.NewConnection(sockAddr),
	}
}
