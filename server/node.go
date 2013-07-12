package server

import (
	"log"
	"net/http"
)

import (
	"gopoker/model"
)

import (
	"code.google.com/p/go.net/websocket"
	"github.com/gorilla/mux"
)

const (
	httpApiRoot  = "/_api"
	wsRoot       = "/_ws"
	rpcRoot      = "/_rpc"
	rpcDebugRoot = "/_rpc/debug"
)

type Node struct {
	Name    string
	ApiAddr string
	//apiService
	RpcAddr string
	//rpcService
	Rooms map[model.Id]*Room
}

func CreateNode(name string, apiAddr string, rpcAddr string) *Node {
	return &Node{
		Name:    name,
		ApiAddr: apiAddr,
		RpcAddr: rpcAddr,
		Rooms:   map[model.Id]*Room{},
	}
}

/*********************************************
********* Rooms
**********************************************/
func (n *Node) Room(id model.Id) *Room {
	room, _ := n.Rooms[id]

	return room
}

func (n *Node) AddRoom(room *Room) bool {
	n.Rooms[room.Id] = room
	return true
}

func (n *Node) RemoveRoom(room *Room) bool {
	delete(n.Rooms, room.Id)
	return true
}

func (n *Node) Start() {
	n.StartRPC()
	n.StartHTTP()
}

func (n *Node) StartHTTP() {
	service := HttpService{n}

	router := mux.NewRouter()

	api := router.PathPrefix(httpApiRoot).Subrouter()

	// Room
	api.HandleFunc("/rooms", service.Rooms).Methods("GET")
	api.HandleFunc("/room/{room}", service.Room).Methods("POST")

	api.HandleFunc("/room/{room}/join", service.Join).Methods("POST")
	api.HandleFunc("/room/{room}/leave", service.Leave).Methods("DELETE")
	api.HandleFunc("/room/{room}/rebuy", service.Rebuy).Methods("POST")
	api.HandleFunc("/room/{room}/addon", service.AddOn).Methods("POST")

	api.HandleFunc("/room/{room}/seating", service.Seating).Methods("GET")
	api.HandleFunc("/room/{room}/wait", service.Wait).Methods("PUT")
	api.HandleFunc("/room/{room}/stats", service.Stats).Methods("GET")

	// misc
	api.HandleFunc("/hand/detect", service.DetectHand).Methods("GET", "POST")
	api.HandleFunc("/hand/random", service.RandomHand).Methods("GET")
	api.HandleFunc("/hand/compare", service.CompareHands).Methods("GET", "POST")
	api.HandleFunc("/hand/odds", service.CalculateOdds).Methods("GET", "POST")

	api.HandleFunc("/deck/generate", service.GenerateDeck).Methods("GET")

	// Deal
	api.HandleFunc("/deal/{deal}", service.Deal).Methods("GET")

	api.HandleFunc("/deal/{deal}/bet", service.Bet).Methods("PUT")
	api.HandleFunc("/deal/{deal}/discard", service.Discard).Methods("PUT")
	api.HandleFunc("/deal/{deal}/muck", service.Muck).Methods("PUT")

	api.HandleFunc("/deal/{deal}/pot", service.Pot).Methods("GET")
	api.HandleFunc("/deal/{deal}/stage", service.Stage).Methods("GET")
	api.HandleFunc("/deal/{deal}/results", service.Results).Methods("GET")
	api.HandleFunc("/deal/{deal}/known_hands", service.KnownHands).Methods("GET")

	// WebSocket
	router.Handle(wsRoot, websocket.Handler(WebSocketHandler))

	log.Printf("starting http service at %s", n.ApiAddr)
	if err := http.ListenAndServe(n.ApiAddr, router); err != nil {
		log.Fatalf("can't start at %s", n.ApiAddr)
	}
}
