package ai

import (
	"net/rpc"
)

import (
	"gopoker/util"
	_"gopoker/client/rpc_client"
)

type Bot struct {
	Id string
	rpcClient *rpc.Client
}

func NewBot(rpcAddr string) *Bot {
	id := util.RandomUuid()
	client, _ := rpc.DialHTTP("tcp", rpcAddr)

	return &Bot{
		Id: id,
		rpcClient: client,
	}
}

func (b *Bot) CreateSession() {

}

func (b *Bot) Join(roomId string, pos int) {

}
