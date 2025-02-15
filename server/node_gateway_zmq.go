package server

import (
	"log"
	_ "time"
)

import (
	"code.google.com/p/goprotobuf/proto"
	zmq "github.com/alecthomas/gozmq"
)

import (
	"gopoker/model"
)

// NodeZMQ - node zeromq service
type NodeZMQ struct {
	*Node

	context   *zmq.Context
	publisher *zmq.Socket
	receiver  *zmq.Socket

	publish     chan [][]byte
	subscribers map[model.Player]chan int
}

// StartZMQ - start zeromq service
func (n *Node) StartZMQ() {
	config := n.Config.ZMQ
	log.Printf("[zmq] starting service, receiver=%s publisher=%s", config.Receiver, config.Publisher)

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

		publish:     make(chan [][]byte, 1000),
		subscribers: map[model.Player]chan int{},
	}

	n.ZMQGateway = gw

	go gw.receive()
	for {
		select {
		case data := <-gw.publish:
			gw.publisher.SendMultipart(data, zmq.NOBLOCK)
		}
	}
}

func (gw *NodeZMQ) receive() {
	for {
		multipart, err := gw.receiver.RecvMultipart(0)
		if err != nil {
			log.Printf("[zmq] PULL error: %s", err)
		}
		if multipart != nil && len(multipart) == 3 {
			log.Printf("[zmq] PULL got multipart: %s", multipart)
			player := model.Player(multipart[0])
			guid := model.Guid(multipart[1])
			if _, subscribed := gw.subscribers[player]; !subscribed {
				go gw.subscribe(player, guid)
			}

			event := &event.Event{}
			if err = event.Unproto(multipart[2]); err != nil {
				log.Printf("[zmq] unproto error: %s", err)
			} else {
				if room, found := gw.Node.Rooms[guid]; found {
					room.Recv <- event
				}
			}
		}
	}
}

func (gw *NodeZMQ) subscribe(player model.Player, guid model.Guid) {
	room := gw.Node.Rooms[guid]

	stop := make(chan int)
	gw.subscribers[player] = stop
	defer delete(gw.subscribers, player)

	recv := make(event.Channel)
	room.Broadcast.Bind(player, &recv)
	defer room.Broadcast.Unbind(player)

Loop:
	for {
		select {
		case event := <-recv:
			log.Printf("[zmq] sending %s to %s", event, player)

			gw.send(event, player)

		case <-stop:
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
