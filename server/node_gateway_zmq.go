package server

import (
	"log"
	"time"
)

import (
	"code.google.com/p/goprotobuf/proto"
	zmq "github.com/alecthomas/gozmq"
)

import (
	"gopoker/event"
	"gopoker/event/message"
	rpc_service "gopoker/server/noderpc"
)

// NodeZMQ - node zeromq service
type NodeZMQ struct {
	*Node

	context   *zmq.Context
	publisher *zmq.Socket
	receiver  *zmq.Socket

	publish chan [][]byte

	connect     chan rpc_service.ConnectGateway
	disconnect  chan rpc_service.DisconnectGateway
	connections map[string]chan int
}

// StartZMQ - start zeromq service
func (n *Node) StartZMQ() {
	config := n.Config.ZMQ
	log.Printf("[zmq] starting service")

	context, _ := zmq.NewContext()
	defer context.Close()

	publisher, _ := context.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Bind(config.Publisher)

	receiver, _ := context.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Bind(config.Receiver)

	gw := &NodeZMQ{
		Node: n,

		context:   context,
		publisher: publisher,
		receiver:  receiver,

		publish: make(chan [][]byte, 1000),

		connect:     make(chan rpc_service.ConnectGateway),
		disconnect:  make(chan rpc_service.DisconnectGateway),
		connections: map[string]chan int{},
	}

	n.ZMQGateway = gw

	go gw.listen()
	go gw.receive()
	gw.accept()
}

func (gw *NodeZMQ) accept() {
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

func (gw *NodeZMQ) listen() {
	for {
		select {
		case data := <-gw.publish:
			gw.publisher.SendMultipart(data, 0)
		}
	}
}

func (gw *NodeZMQ) receive() {
	for {
		for data, _ := gw.receiver.RecvMultipart(zmq.NOBLOCK); data != nil; {
			msg := &message.Message{}

			roomID := string(data[0])

			err := proto.Unmarshal(data[1], msg)
			if err != nil {
				log.Printf("[zmq] unmarshal error: %s", err)
			} else {
				room := gw.Node.Rooms[roomID]
				room.Recv <- msg
			}
		}
		time.Sleep(1e6)
	}
}

func (gw *NodeZMQ) startConnection(playerID string, roomID string, stop *chan int) {
	recv := make(event.MessageChannel, 100)

	room := gw.Node.Rooms[roomID]

	room.Broadcast.Broker.Bind(event.Player, playerID, &recv)
	defer room.Broadcast.Broker.Unbind(event.Player, playerID)

Loop:
	for {
		select {
		case msg := <-recv:
			//log.Printf("[zmq] sending %s to %s", msg, playerID)
			gw.send(msg, playerID)

		case <-*stop:
			log.Printf("[zmq] stop connection for %s", playerID)

			break Loop
		}
	}
}

func (gw *NodeZMQ) send(msg *message.Message, playerID string) {
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Printf("[zmq] marshal error: %s", err)
	} else {
		gw.publish <- [][]byte{[]byte(playerID), data}
	}
}
