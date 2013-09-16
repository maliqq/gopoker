package server

import (
	stomp_server "github.com/jjeffery/stomp/server"
)

type NodeStomp struct {
	server stomp_server.Server
}
