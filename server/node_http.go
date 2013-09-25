package server

import (
	"encoding/json"
	"log"
	"net/http"
)

import (
	"code.google.com/p/go.net/websocket"
	gorilla_mux "github.com/gorilla/mux"
	gorilla_rpc "github.com/gorilla/rpc"
	gorilla_json "github.com/gorilla/rpc/json"
)

// NodeHTTP - http server
type NodeHTTP struct {
	*Node
}

// Rooms - mapping by id
type Rooms struct {
	Rooms map[string]*Room
}

// StartHTTP - starts http service
func (n *Node) StartHTTP() {
	router := gorilla_mux.NewRouter()
	n.drawRoutes(router)

	addr := n.Config.HTTP.Addr
	log.Printf("[http] starting service at %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("[http] can't start at %s", addr)
	}
}

func (n *Node) drawRoutes(router *gorilla_mux.Router) {
	config := n.Config.HTTP

	// REST API
	nodeHTTP := &NodeHTTP{n}
	api := router.PathPrefix(config.GetAPIPath()).Subrouter()
	nodeHTTP.drawAPI(api)

	// JSON-RPC over HTTP
	nodeRPC := NodeRPC{n}
	rpc := gorilla_rpc.NewServer()
	err := rpc.RegisterService(nodeRPC, "")
	if err != nil {
		log.Fatalf("[http-rpc] %s", err)
	}
	rpc.RegisterCodec(gorilla_json.NewCodec(), "application/json")

	// handle RPC
	router.HandleFunc(config.GetRPCPath(), func(resp http.ResponseWriter, req *http.Request) {
		RespondCORS(resp)
		resp.Write([]byte{0xA})
	}).Methods("OPTIONS")

	router.HandleFunc(config.GetRPCPath(), func(resp http.ResponseWriter, req *http.Request) {
		RespondCORS(resp)
		rpc.ServeHTTP(resp, req)
	}).Methods("POST")

	// handle WebSocket
	router.Handle(config.GetWebSocketPath(), websocket.Handler(nodeHTTP.WebSocketHandler))
}

func (nodeHTTP *NodeHTTP) drawAPI(api *gorilla_mux.Router) {
	// Room
	api.HandleFunc("/rooms", nodeHTTP.Rooms).Methods("GET", "OPTIONS")
	api.HandleFunc("/room/{id}", nodeHTTP.Room).Methods("GET", "OPTIONS")

	api.HandleFunc("/room/{id}/join", nodeHTTP.Join).Methods("POST")
	api.HandleFunc("/room/{id}/leave", nodeHTTP.Leave).Methods("DELETE")
	api.HandleFunc("/room/{id}/rebuy", nodeHTTP.Rebuy).Methods("POST")
	api.HandleFunc("/room/{id}/addon", nodeHTTP.AddOn).Methods("POST")

	api.HandleFunc("/room/{room}/seating", nodeHTTP.Seating).Methods("GET")
	api.HandleFunc("/room/{room}/wait", nodeHTTP.Wait).Methods("PUT")
	api.HandleFunc("/room/{room}/stats", nodeHTTP.Stats).Methods("GET")

	// misc
	api.HandleFunc("/hand/detect", nodeHTTP.DetectHand).Methods("GET", "POST")
	api.HandleFunc("/hand/random", nodeHTTP.RandomHand).Methods("GET")
	api.HandleFunc("/hand/compare", nodeHTTP.CompareHands).Methods("GET", "POST")
	api.HandleFunc("/hand/odds", nodeHTTP.CalculateOdds).Methods("GET", "POST")

	api.HandleFunc("/deck/generate", nodeHTTP.GenerateDeck).Methods("GET")

	// Play
	api.HandleFunc("/play/{id}", nodeHTTP.Play).Methods("GET")

	api.HandleFunc("/play/{id}/bet", nodeHTTP.Bet).Methods("PUT")
	api.HandleFunc("/play/{id}/discard", nodeHTTP.Discard).Methods("PUT")
	api.HandleFunc("/play/{id}/muck", nodeHTTP.Muck).Methods("PUT")

	api.HandleFunc("/play/{id}/pot", nodeHTTP.Pot).Methods("GET")
	api.HandleFunc("/play/{id}/stage", nodeHTTP.Stage).Methods("GET")
	api.HandleFunc("/play/{id}/winners", nodeHTTP.Winners).Methods("GET")
	api.HandleFunc("/play/{id}/known_cards", nodeHTTP.KnownCards).Methods("GET")
}

// Log - log request
func (nodeHTTP *NodeHTTP) Log(req *http.Request) {
	// nginx default format:
	//$remote_addr - $remote_user [$time_local]  "$request" $status $bytes_sent "$http_referer" "$http_user_agent" "$gzip_ratio"
	log.Printf("%s - [%s %s %s] %s\n", req.RemoteAddr, req.Method, req.RequestURI, req.Proto, req.UserAgent())
}

// RespondCORS - respond with CORS headers
func RespondCORS(resp http.ResponseWriter) {
	// CORS headers
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type")
}

// RespondJSON - respond with JSON
func (nodeHTTP *NodeHTTP) RespondJSON(resp http.ResponseWriter, result interface{}) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("[http] Can't marshal object: %+v", err)
	}

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	// CORS headers
	RespondCORS(resp)

	resp.Write(data)
	resp.Write([]byte{0xA})
}

// RespondJSONError - respond with JSON error
func (nodeHTTP *NodeHTTP) RespondJSONError(resp http.ResponseWriter, err error) {
	data := struct {
		Error string
	}{
		Error: err.Error(),
	}

	nodeHTTP.RespondJSON(resp, data)
}
