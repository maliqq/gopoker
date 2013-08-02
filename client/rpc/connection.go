package rpc

import (
	"net/rpc"
)

// Connection - RPC tcp connection
type Connection struct {
	*rpc.Client
}

// Send - write data
func (c *Connection) Send(data interface{}) error {
	return nil
}

// Receive - read data
func (c *Connection) Receive(data interface{}) error {
	return nil
}
