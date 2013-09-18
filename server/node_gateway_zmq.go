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
	"gopoker/model"
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
	connections map[model.Player]chan int
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
	receiver.Connect(config.Receiver)

	gw := &NodeZMQ{
		Node: n,

		context:   context,
		publisher: publisher,
		receiver:  receiver,

		publish: make(chan [][]byte, 1000),

		connect:     make(chan rpc_service.ConnectGateway),
		disconnect:  make(chan rpc_service.DisconnectGateway),
		connections: map[model.Player]chan int{},
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
			gw.connections[req.Player] = stop
			go gw.startConnection(req.Player, req.Room, &stop)

		case req := <-gw.disconnect:
			log.Printf("[zmq] disconnect request: %+v", req)
			player := model.Player(req.Player)
			stop, ok := gw.connections[player]
			if ok {
				stop <- 1
				delete(gw.connections, player)
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
		for data, err := gw.receiver.RecvMultipart(zmq.NOBLOCK); err == nil; {
			event := &event.Event{}
			guid := string(data[0])

			err := event.UnmarshalProto(data[1])
			if err != nil {
				log.Printf("[zmq] unmarshal error: %s", err)
			} else {
				room := gw.Node.Rooms[model.Guid(guid)]
				room.Recv <- event
			}
		}
		time.Sleep(1e6)
	}
}

func (gw *NodeZMQ) startConnection(player model.Player, guid model.Guid, stop *chan int) {
	recv := make(event.Channel, 100)

	room := gw.Node.Rooms[guid]

	room.Broadcast.Bind(player, &recv)
	defer room.Broadcast.Unbind(player)

Loop:
	for {
		select {
		case event := <-recv:
			//log.Printf("[zmq] sending %s to %s", msg, playerID)
			gw.send(event, player)

		case <-*stop:
			log.Printf("[zmq] stop connection for %s", player)

			break Loop
		}
	}
}

func (gw *NodeZMQ) send(event *event.Event, player model.Player) {
	data, err := proto.Marshal(event.Proto())
	if err != nil {
		log.Printf("[zmq] marshal error: %s", err)
	} else {
		gw.publish <- [][]byte{[]byte(player), data}
	}
}
