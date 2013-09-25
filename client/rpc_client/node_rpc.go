package rpc_client

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

import (
	"gopoker/event"
	"gopoker/message"
	"gopoker/model"
)

// Connection - RPC tcp connection
type NodeRPC struct {
	*rpc.Client
}

func NewConnection(rpcAddr string) *NodeRPC {
	client, err := jsonrpc.Dial("tcp", rpcAddr)
	if err != nil {
		log.Fatal("dialing error: ", err)
	}
	return &NodeRPC{
		Client: client,
	}
}

func (c *NodeRPC) CreateRoom(guid model.Guid, betSize float64, variation model.Variation) event.CallResult {
	msg := message.CreateRoom{
		Guid: guid,
		BetSize: betSize,
	}
	
	switch g := variation.(type) {
	case *model.Mix:
		msg.Mix = g
	case *model.Game:
		msg.Game = g
	}

	return c.Call("CreateRoom", msg)
}

func (c *NodeRPC) Call(method string, args interface{}) event.CallResult {
	var result event.CallResult

	err := c.Client.Call("NodeRPC." + method, args, &result)

	if err != nil {
		log.Fatal("rpc call error: ", err)
	}

	return result
}
