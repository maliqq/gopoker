package server

import (
	"code.google.com/p/goprotobuf/proto"
	zmq "github.com/alecthomas/gozmq"
)

import (
	"gopoker/protocol"
	rpc_service "gopoker/server/noderpc"
)

// NodeZMQ - node zeromq service
type NodeZMQ struct {
	*Node
	context     *zmq.Context
	socket      *zmq.Socket
	connect     chan rpc_service.ConnectGateway
	disconnect  chan rpc_service.DisconnectGateway
	connections map[string]chan int
}

// StartZMQ - start zeromq service
func (n *Node) StartZMQ() {
	addr := n.Config.ZMQ

	context, _ := zmq.NewContext()
	defer context.Close()

	socket, _ := context.NewSocket(zmq.PUB)
	defer socket.Close()

	socket.Bind(addr)

	nodeZMQ := &NodeZMQ{
		Node:    n,
		context: context,
		socket:  socket,
	}

	n.ZMQGateway = nodeZMQ

	for {
		select {
		case req := <-nodeZMQ.connect:
			stop := make(chan int)
			nodeZMQ.connections[req.PlayerID] = stop
			go nodeZMQ.startConnection(req.PlayerID, req.RoomID, &stop)

		case req := <-nodeZMQ.disconnect:
			stop, ok := nodeZMQ.connections[req.PlayerID]
			if ok {
				stop <- 1
				delete(nodeZMQ.connections, req.PlayerID)
			}
		}
	}
}

func (nodeZMQ *NodeZMQ) startConnection(playerID string, roomID string, stop *chan int) {
	recv := make(protocol.MessageChannel)

	room := nodeZMQ.Node.Rooms[roomID]
	room.Broadcast.BindUser(playerID, &recv)

Loop:
	for {
		select {
		case msg := <-recv:
			data, err := proto.Marshal(msg)
			if err != nil {
				nodeZMQ.socket.Send(data, 0)
			}
		case <-*stop:
			break Loop
		}
	}
}
