package server

import (
	"encoding/json"
	"log"
	"net/http"
)

import (
	"code.google.com/p/go.net/websocket"
	"github.com/gorilla/mux"
)

import (
	"gopoker/model"
)

const (
	WebSocketPath = "/_ws"
	HttpApiPath   = "/_api"
)

type NodeHTTP struct {
	*Node
}

type Rooms struct {
	Rooms map[model.Id]*Room
}

func (n *Node) StartHTTP() {
	nodeHTTP := &NodeHTTP{n}

	router := mux.NewRouter()
	nodeHTTP.drawRoutes(router)

	log.Printf("[http] Starting service at %s", n.ApiAddr)
	if err := http.ListenAndServe(n.ApiAddr, router); err != nil {
		log.Fatalf("[http] Can't start at %s", n.ApiAddr)
	}
}

func (nodeHTTP *NodeHTTP) drawRoutes(router *mux.Router) {
	api := router.PathPrefix(HttpApiPath).Subrouter()

	// Room
	api.HandleFunc("/rooms", nodeHTTP.Rooms).Methods("GET", "OPTIONS")
	api.HandleFunc("/room/{room}", nodeHTTP.Room).Methods("GET", "OPTIONS")

	api.HandleFunc("/room/{room}/join", nodeHTTP.Join).Methods("POST")
	api.HandleFunc("/room/{room}/leave", nodeHTTP.Leave).Methods("DELETE")
	api.HandleFunc("/room/{room}/rebuy", nodeHTTP.Rebuy).Methods("POST")
	api.HandleFunc("/room/{room}/addon", nodeHTTP.AddOn).Methods("POST")

	api.HandleFunc("/room/{room}/seating", nodeHTTP.Seating).Methods("GET")
	api.HandleFunc("/room/{room}/wait", nodeHTTP.Wait).Methods("PUT")
	api.HandleFunc("/room/{room}/stats", nodeHTTP.Stats).Methods("GET")

	// misc
	api.HandleFunc("/hand/detect", nodeHTTP.DetectHand).Methods("GET", "POST")
	api.HandleFunc("/hand/random", nodeHTTP.RandomHand).Methods("GET")
	api.HandleFunc("/hand/compare", nodeHTTP.CompareHands).Methods("GET", "POST")
	api.HandleFunc("/hand/odds", nodeHTTP.CalculateOdds).Methods("GET", "POST")

	api.HandleFunc("/deck/generate", nodeHTTP.GenerateDeck).Methods("GET")

	// Deal
	api.HandleFunc("/deal/{deal}", nodeHTTP.Deal).Methods("GET")

	api.HandleFunc("/deal/{deal}/bet", nodeHTTP.Bet).Methods("PUT")
	api.HandleFunc("/deal/{deal}/discard", nodeHTTP.Discard).Methods("PUT")
	api.HandleFunc("/deal/{deal}/muck", nodeHTTP.Muck).Methods("PUT")

	api.HandleFunc("/deal/{deal}/pot", nodeHTTP.Pot).Methods("GET")
	api.HandleFunc("/deal/{deal}/stage", nodeHTTP.Stage).Methods("GET")
	api.HandleFunc("/deal/{deal}/results", nodeHTTP.Results).Methods("GET")
	api.HandleFunc("/deal/{deal}/known_hands", nodeHTTP.KnownHands).Methods("GET")

	// WebSocket
	router.Handle(WebSocketPath, websocket.Handler(nodeHTTP.WebSocketHandler))
}

func (nodeHTTP *NodeHTTP) Log(req *http.Request) {
	// nginx default format:
	//$remote_addr - $remote_user [$time_local]  "$request" $status $bytes_sent "$http_referer" "$http_user_agent" "$gzip_ratio"
	log.Printf("%s - [%s %s %s] %s\n", req.RemoteAddr, req.Method, req.RequestURI, req.Proto, req.UserAgent())
}

func (nodeHTTP *NodeHTTP) RespondJSON(resp http.ResponseWriter, result interface{}) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("[http] Can't marshal object: %+v", result)
	}

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	// CORS headers
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")

	resp.Write(data)
	resp.Write([]byte{0xA})
}
