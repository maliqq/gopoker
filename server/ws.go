package server

import (
	"code.google.com/p/go.net/websocket"
)

import (
	"gopoker/client"
	"gopoker/client/websocket_client"
	"gopoker/model"
	"gopoker/util"
)

func (nodeHTTP *NodeHTTP) WebSocketHandler(conn *websocket.Conn) {
	id := model.Id(util.RandomUuid())
	connection := &websocket_client.Connection{conn}
	session := client.NewSession(connection)

	nodeHTTP.Node.Sessions[id] = session

	defer func() {
		delete(nodeHTTP.Node.Sessions, id)
		close(session.Receive)
	}()

	session.Start()
}
