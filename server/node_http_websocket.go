package server

import (
	"fmt"
)

import (
	"code.google.com/p/go.net/websocket"
)

import (
	"gopoker/client/websocket_client"
	"gopoker/event/message"
	"gopoker/model"
)

// WebSocketHandler - websocket connection handler
func (nodeHTTP *NodeHTTP) WebSocketHandler(conn *websocket.Conn) {
	node := nodeHTTP.Node
	q := conn.Request().URL.Query()
	roomID := q.Get("room_id")
	room, found := node.Rooms[model.Guid(roomID)]

	if !found {
		errMsg := message.ErrorMessage{fmt.Errorf("room with id=%s not found", roomID)}
		websocket.JSON.Send(conn, errMsg)
		conn.Close()
		return
	}

	id := NewSessionID()
	connection := &websocket_client.Connection{Conn: conn}
	session := NewSession(id, connection)
	session.Send = &room.Recv

	//session.Connection.Send(room)

	//for _, player := range room.Table.AllPlayers() {
	//	room.Broadcast.Bind(player, &session.Recv)
	//}
	room.Broadcast.Bind(id, &session.Recv)
	defer func() {
		room.Broadcast.Unbind(id)
		session.Close()
	}()

	session.Start()
}
