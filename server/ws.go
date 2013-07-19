package server

import (
	"code.google.com/p/go.net/websocket"
)

import (
	"gopoker/client"
	"gopoker/client/websocket_client"
	"gopoker/protocol"
	"gopoker/util"
)

func (nodeHTTP *NodeHTTP) WebSocketHandler(conn *websocket.Conn) {
	node := nodeHTTP.Node
	q := conn.Request().URL.Query()
	roomId := q.Get("room_id")
	_, found := node.Rooms[roomId]

	if !found {
		// 404
	}

	id := util.RandomUuid()
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
