package server

import (
	"fmt"
)

import (
	"code.google.com/p/go.net/websocket"
)

import (
	"gopoker/client"
	"gopoker/client/websocket_client"
	"gopoker/model"
	"gopoker/protocol"
	"gopoker/util"
)

func (nodeHTTP *NodeHTTP) WebSocketHandler(conn *websocket.Conn) {
	node := nodeHTTP.Node
	q := conn.Request().URL.Query()
	roomId := model.Id(q.Get("room_id"))
	room := node.Rooms[roomId]

	id := model.Id(util.RandomUuid())
	connection := &websocket_client.Connection{conn}
	session := client.NewSession(connection)
	//session.Connection.Send(room)

	me := make(protocol.MessageChannel)

	nodeHTTP.Node.Sessions[id] = session

	defer func() {
		delete(nodeHTTP.Node.Sessions, id)
		close(me)
	}()

	session.Start(&me)
}
