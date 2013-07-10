package server

import (
	"code.google.com/p/go.net/websocket"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net"
	"net/http"
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
	Rooms map[string]*Room
}

func CreateNode(name string, apiAddr string, rpcAddr string) *Node {
	return &Node{
		Name:    name,
		ApiAddr: apiAddr,
		RpcAddr: rpcAddr,
		Rooms:   map[string]*Room{},
	}
}

func (n *Node) startRpcService() {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterService(n, "")

	http.Handle(rpcRoot, rpcServer)

	log.Printf("starting rpc service at %s", n.RpcAddr)
	socket, err := net.Listen("tcp", n.RpcAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(socket, nil)
}

func (n *Node) startHttpService() {
	service := HttpService{node: n}

	router := mux.NewRouter()

	api := router.PathPrefix(httpApiRoot).Subrouter()

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

	// Table
	api.HandleFunc("/table/{table}", service.Table).Methods("POST")

	api.HandleFunc("/table/{table}/join", service.JoinTable).Methods("POST")
	api.HandleFunc("/table/{table}/leave", service.LeaveTable).Methods("DELETE")
	api.HandleFunc("/table/{table}/rebuy", service.Rebuy).Methods("POST")
	api.HandleFunc("/table/{table}/addon", service.AddOn).Methods("POST")

	api.HandleFunc("/table/{table}/seating", service.TableSeating).Methods("GET")
	api.HandleFunc("/table/{table}/wait", service.Wait).Methods("PUT")
	api.HandleFunc("/table/{table}/stats", service.TableStats).Methods("GET")

	// WebSocket
	router.Handle(wsRoot, websocket.Handler(WebSocketHandler))

	log.Printf("starting http service at %s", n.ApiAddr)
	err := http.ListenAndServe(n.ApiAddr, router)
	if err != nil {
		log.Fatalf("can't start at %s", n.ApiAddr)
	}
}

func (n *Node) startStatsWorker() {

}

func (n *Node) Start() {
	//n.startRpcService()
	n.startHttpService()
	//go n.startStatsWorker()
}
