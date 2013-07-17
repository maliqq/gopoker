package rpc_client

import (
	"net/rpc"
)

type Connection struct {
	*rpc.Client
}

func (c *Connection) Send(data interface{}) error {
	return nil
}

func (c *Connection) Receive(data interface{}) error {
	return nil
}
