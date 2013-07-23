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
	"gopoker/protocol"
	"gopoker/util"
)

func (nodeHTTP *NodeHTTP) WebSocketHandler(conn *websocket.Conn) {
	node := nodeHTTP.Node
	q := conn.Request().URL.Query()
	roomId := q.Get("room_id")
	room, found := node.Rooms[roomId]

	if !found {
		errMsg := protocol.NewError(fmt.Errorf("room with id=%s not found", roomId))
		websocket.JSON.Send(conn, errMsg)
		conn.Close()
		return
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
