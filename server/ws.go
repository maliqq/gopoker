package server

import (
	"code.google.com/p/go.net/websocket"
)

import (
	"gopoker/client/ws"
	"gopoker/model"
	"gopoker/util"
)

func (nodeHTTP *NodeHTTP) WebSocketHandler(connection *websocket.Conn) {
	id := model.Id(util.RandomUuid())
	session := ws.NewSession(connection)

	nodeHTTP.Node.Sessions[id] = session

	defer func() {
		delete(nodeHTTP.Node.Sessions, id)
		close(session.Receive)
	}()

	session.Start()
}
