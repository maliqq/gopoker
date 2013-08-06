package rpc

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

import (
	rpc_service "gopoker/server/noderpc"
)

// Connection - RPC tcp connection
type Connection struct {
	*rpc.Client
}

func NewConnection(rpcAddr string) *Connection {
	client, err := jsonrpc.Dial("tcp", rpcAddr)
	if err != nil {
		log.Fatal("dialing error: ", err)
	}
	return &Connection{
		Client: client,
	}
}

func (c *Connection) Call(method string, args interface{}) rpc_service.CallResult {
	var result rpc_service.CallResult

	err := c.Client.Call(fmt.Sprintf("NodeRPC.%s", method), args, &result)

	if err != nil {
		log.Fatal("rpc call error: ", err)
	}

	return result
}
