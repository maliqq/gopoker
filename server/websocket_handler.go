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
	room, found := node.Rooms[roomId]

	if !found {
		// 404
	}

	id := util.RandomUuid()
	connection := &websocket_client.Connection{conn}
	session := client.NewSession(id, connection)

	//session.Connection.Send(room)

	session.Receive = make(protocol.MessageChannel)
	session.Send = &room.Receive
	defer close(session.Receive)

	//for _, player := range room.Table.AllPlayers() {
	//	room.Broadcast.Bind(player, &session.Receive)
	//}
	room.Broadcast.Bind(protocol.Private, &session.Receive)
	room.Broadcast.Bind(protocol.Public, &session.Receive)

	session.Start()
}
