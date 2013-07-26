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

const (
	DefaultApiPath       = "/_api"
	DefaultRpcPath       = "/_rpc"
	DefaultWebSocketPath = "/_ws"
)

type NodeHTTP struct {
	*Node
}

type Rooms struct {
	Rooms map[string]*Room
}

func (n *Node) StartHTTP() {
	router := gorilla_mux.NewRouter()
	n.drawRoutes(router)

	log.Printf("[http] Starting service at %s", n.Http.Addr)
	if err := http.ListenAndServe(n.Http.Addr, router); err != nil {
		log.Fatalf("[http] Can't start at %s", n.Http.Addr)
	}
}

func paths(httpConfig *HttpConfig) (string, string, string) {
	apiPath := httpConfig.ApiPath
	if apiPath == "" {
		apiPath = DefaultApiPath
	}

	webSocketPath := httpConfig.WebSocketPath
	if webSocketPath == "" {
		webSocketPath = DefaultWebSocketPath
	}

	rpcPath := httpConfig.RpcPath
	if rpcPath == "" {
		rpcPath = DefaultRpcPath
	}

	return apiPath, rpcPath, webSocketPath
}

func (n *Node) drawRoutes(router *gorilla_mux.Router) {
	apiPath, rpcPath, webSocketPath := paths(n.Config.Http)

	// REST API
	nodeHTTP := &NodeHTTP{n}
	api := router.PathPrefix(apiPath).Subrouter()
	nodeHTTP.drawApi(api)

	// JSON-RPC over HTTP
	nodeRPC := NodeRPC{n}
	rpc := gorilla_rpc.NewServer()
	rpc.RegisterService(nodeRPC, "")
	rpc.RegisterCodec(gorilla_json.NewCodec(), "application/json")

	// handle RPC
	router.Handle(rpcPath, rpc)

	// handle WebSocket
	router.Handle(webSocketPath, websocket.Handler(nodeHTTP.WebSocketHandler))
}

func (nodeHTTP *NodeHTTP) drawApi(api *gorilla_mux.Router) {
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
