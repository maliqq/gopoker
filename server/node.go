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
	api.HandleFunc("/hand/detect", func(resp http.ResponseWriter, req *http.Request) {
		service.DetectHand(resp, req)
	}).Methods("GET", "POST")
	api.HandleFunc("/hand/random", func(resp http.ResponseWriter, req *http.Request) {
		service.RandomHand(resp, req)
	}).Methods("GET")
	api.HandleFunc("/hand/compare", func(resp http.ResponseWriter, req *http.Request) {
		service.CompareHands(resp, req)
	}).Methods("GET", "POST")
	api.HandleFunc("/hand/odds", func(resp http.ResponseWriter, req *http.Request) {
		service.CalculateOdds(resp, req)
	}).Methods("GET", "POST")

	api.HandleFunc("/deck/generate", func(resp http.ResponseWriter, req *http.Request) {
		service.GenerateDeck(resp, req)
	}).Methods("GET")

	// Deal
	api.HandleFunc("/deal/{deal}", func(resp http.ResponseWriter, req *http.Request) {
		service.GetDeal(resp, req)
	}).Methods("GET")

	api.HandleFunc("/deal/{deal}/bet", func(resp http.ResponseWriter, req *http.Request) {
		service.Bet(resp, req)
	}).Methods("PUT")
	api.HandleFunc("/deal/{deal}/discard", func(resp http.ResponseWriter, req *http.Request) {
		service.Discard(resp, req)
	}).Methods("PUT")
	api.HandleFunc("/deal/{deal}/muck", func(resp http.ResponseWriter, req *http.Request) {
		service.Muck(resp, req)
	}).Methods("PUT")

	api.HandleFunc("/deal/{deal}/pot", func(resp http.ResponseWriter, req *http.Request) {
		service.GetPot(resp, req)
	}).Methods("GET")
	api.HandleFunc("/deal/{deal}/stage", func(resp http.ResponseWriter, req *http.Request) {
		service.GetStage(resp, req)
	}).Methods("GET")
	api.HandleFunc("/deal/{deal}/results", func(resp http.ResponseWriter, req *http.Request) {
		service.GetResults(resp, req)
	}).Methods("GET")
	api.HandleFunc("/deal/{deal}/known_hands", func(resp http.ResponseWriter, req *http.Request) {
		service.GetKnownHands(resp, req)
	}).Methods("GET")

	// Table
	api.HandleFunc("/table/{table}", func(resp http.ResponseWriter, req *http.Request) {
		service.GetTable(resp, req)
	}).Methods("POST")

	api.HandleFunc("/table/{table}/join", func(resp http.ResponseWriter, req *http.Request) {
		service.JoinTable(resp, req)
	}).Methods("POST")
	api.HandleFunc("/table/{table}/leave", func(resp http.ResponseWriter, req *http.Request) {
		service.LeaveTable(resp, req)
	}).Methods("DELETE")
	api.HandleFunc("/table/{table}/rebuy", func(resp http.ResponseWriter, req *http.Request) {
		service.Rebuy(resp, req)
	}).Methods("POST")
	api.HandleFunc("/table/{table}/addon", func(resp http.ResponseWriter, req *http.Request) {
		service.AddOn(resp, req)
	}).Methods("POST")

	api.HandleFunc("/table/{table}/seating", func(resp http.ResponseWriter, req *http.Request) {
		service.GetTableSeating(resp, req)
	}).Methods("GET")
	api.HandleFunc("/table/{table}/wait", func(resp http.ResponseWriter, req *http.Request) {
		service.Wait(resp, req)
	}).Methods("PUT")
	api.HandleFunc("/table/{table}/stats", func(resp http.ResponseWriter, req *http.Request) {
		service.GetTableStats(resp, req)
	}).Methods("GET")

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
