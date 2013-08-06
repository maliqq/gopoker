package server

import (
	"log"
	"sync"
)

import (
	"code.google.com/p/goprotobuf/proto"
	zmq "github.com/alecthomas/gozmq"
)

import (
	"gopoker/protocol"
	"gopoker/protocol/message"
	rpc_service "gopoker/server/noderpc"
)

// NodeZMQ - node zeromq service
type NodeZMQ struct {
	*Node
	context     *zmq.Context
	socket      *zmq.Socket
	sendlock    sync.Mutex
	connect     chan rpc_service.ConnectGateway
	disconnect  chan rpc_service.DisconnectGateway
	connections map[string]chan int
}

// StartZMQ - start zeromq service
func (n *Node) StartZMQ() {
	addr := n.Config.ZMQ
	log.Printf("[zmq] starting service at %s", addr)

	context, _ := zmq.NewContext()
	defer context.Close()

	socket, _ := context.NewSocket(zmq.PUB)
	defer socket.Close()

	socket.Bind(addr)

	gw := &NodeZMQ{
		Node:        n,
		context:     context,
		socket:      socket,
		connect:     make(chan rpc_service.ConnectGateway),
		disconnect:  make(chan rpc_service.DisconnectGateway),
		connections: map[string]chan int{},
	}

	n.ZMQGateway = gw

	for {
		select {
		case req := <-gw.connect:
			log.Printf("[zmq] connect request: %+v", req)
			stop := make(chan int)
			gw.connections[req.PlayerID] = stop
			go gw.startConnection(req.PlayerID, req.RoomID, &stop)

		case req := <-gw.disconnect:
			log.Printf("[zmq] disconnect request: %+v", req)
			stop, ok := gw.connections[req.PlayerID]
			if ok {
				stop <- 1
				delete(gw.connections, req.PlayerID)
			}
		}
	}
}

func (gw *NodeZMQ) startConnection(playerID string, roomID string, stop *chan int) {
	recv := make(protocol.MessageChannel, 100)

	room := gw.Node.Rooms[roomID]
	room.Broadcast.BindUser(playerID, &recv)

Loop:
	for {
		select {
		case msg := <-recv:
			//log.Printf("[zmq] sending %s to %s", msg, playerID)
			go gw.send(msg, playerID)

		case <-*stop:
			log.Printf("[zmq] stop connection for %s", playerID)

			break Loop
		}
	}

	room.Broadcast.UnbindUser(playerID)
}

func (gw *NodeZMQ) send(msg *message.Message, playerID string) {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Printf("[zmq] marshal error: %s", err)
	} else {
		gw.sendlock.Lock()
		gw.socket.SendMultipart([][]byte{[]byte(playerID), data}, 0)
		gw.sendlock.Unlock()
	}
}
