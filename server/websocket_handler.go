package server

import (
	"fmt"
)

import (
	"code.google.com/p/go.net/websocket"
)

import (
	websocket_client "gopoker/client/ws"
	"gopoker/protocol"
	"gopoker/protocol/message"
	"gopoker/util"
)

// WebSocketHandler - websocket connection handler
func (nodeHTTP *NodeHTTP) WebSocketHandler(conn *websocket.Conn) {
	node := nodeHTTP.Node
	q := conn.Request().URL.Query()
	roomID := q.Get("room_id")
	room, found := node.Rooms[roomID]

	if !found {
		errMsg := message.NotifyErrorMessage(fmt.Errorf("room with id=%s not found", roomID))
		websocket.JSON.Send(conn, errMsg)
		conn.Close()
		return
	}

	id := util.RandomUuid()
	connection := &websocket_client.Connection{Conn: conn}
	session := NewSession(id, connection)
	session.Send = &room.Recv

	//session.Connection.Send(room)

	//for _, player := range room.Table.AllPlayers() {
	//	room.Broadcast.Bind(player, &session.Recv)
	//}
	room.Broadcast.Broker.Bind(protocol.Watcher, "test", &session.Recv)
	defer func() {
		room.Broadcast.Broker.Unbind(protocol.Watcher, "test")
		session.Close()
	}()

	session.Start()
}
