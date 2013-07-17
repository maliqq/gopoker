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
	"gopoker/poker"
	"gopoker/poker/ranking"
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

	log.Printf("starting HTTP service at %s", n.ApiAddr)
	if err := http.ListenAndServe(n.ApiAddr, router); err != nil {
		log.Fatalf("can't start at %s", n.ApiAddr)
	}
}

func (nodeHTTP *NodeHTTP) drawRoutes(router *mux.Router) {
	api := router.PathPrefix(HttpApiPath).Subrouter()

	// Room
	api.HandleFunc("/rooms", nodeHTTP.Rooms).Methods("GET")
	api.HandleFunc("/room/{room}", nodeHTTP.Room).Methods("POST")

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

func (nodeHTTP *NodeHTTP) Rooms(resp http.ResponseWriter, req *http.Request) {
	rooms := nodeHTTP.Node.Rooms

	nodeHTTP.RespondJSON(resp, rooms)
}

func (nodeHTTP *NodeHTTP) Room(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	id := q.Get("room")
	room := nodeHTTP.Node.Room(model.Id(id))

	nodeHTTP.RespondJSON(resp, room)
}

func (nodeHTTP *NodeHTTP) Log(req *http.Request) {
	// nginx default format:
	//$remote_addr - $remote_user [$time_local]  "$request" $status $bytes_sent "$http_referer" "$http_user_agent" "$gzip_ratio"
	log.Printf("%s - [%s %s %s] %s\n", req.RemoteAddr, req.Method, req.RequestURI, req.Proto, req.UserAgent())
}

func (nodeHTTP *NodeHTTP) RespondJSON(resp http.ResponseWriter, result interface{}) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("Can't marshal object: %+v", result)
		return
	}

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.Header().Set("Access-Control-Allow-Origin", "*")

	resp.Write(data)
	resp.Write([]byte{0xA})
}

type CompareResult struct {
	A      *poker.Hand
	B      *poker.Hand
	Board  *poker.Cards
	Result int
}

type OddsResult struct {
	A     *poker.Cards
	B     *poker.Cards
	Total int
	Wins  float64
	Ties  float64
	Loses float64
}

type pocketHand struct {
	Pocket poker.Cards
	Hand   *poker.Hand
}

type dealHand struct {
	Board   poker.Cards
	Pockets []pocketHand
}

func (nodeHTTP *NodeHTTP) DetectHand(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	r := q.Get("ranking")
	if r == "" {
		r = "high"
	}
	ranking := ranking.Ranking(r)

	if c := q.Get("cards"); c != "" {
		cards, err := poker.ParseCards(c)
		if err != nil {
			resp.Write([]byte(err.Error()))
			return
		}

		hand, err := poker.Detect[ranking](cards)
		if err != nil {
			resp.Write([]byte(err.Error()))
			return
		}

		nodeHTTP.RespondJSON(resp, hand)
	} else {
		resp.Write([]byte("no cards specified"))
	}
}

func (nodeHTTP *NodeHTTP) CompareHands(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	a, _ := poker.ParseCards(q.Get("a"))
	b, _ := poker.ParseCards(q.Get("b"))

	dealer := model.NewDealer()
	dealer.Burn(a)
	dealer.Burn(b)
	board := dealer.Share(5)
	c1 := append(*a, *board...)
	c2 := append(*b, *board...)
	h1, _ := poker.Detect[ranking.High](&c1)
	h2, _ := poker.Detect[ranking.High](&c2)

	s, _ := json.Marshal(&CompareResult{
		A:      h1,
		B:      h2,
		Board:  board,
		Result: h1.Compare(h2),
	})

	resp.Write([]byte(s))
}

func (nodeHTTP *NodeHTTP) CalculateOdds(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	a, _ := poker.ParseCards(q.Get("a"))
	b, _ := poker.ParseCards(q.Get("b"))

	total := 10000
	wins, ties, loses := 0, 0, 0
	for i := 0; i <= total; i++ {
		dealer := model.NewDealer()
		dealer.Burn(a)
		dealer.Burn(b)
		board := dealer.Share(5)
		c1 := append(*a, *board...)
		c2 := append(*b, *board...)
		h1, _ := poker.Detect[ranking.High](&c1)
		h2, _ := poker.Detect[ranking.High](&c2)

		switch h1.Compare(h2) {
		case -1:
			loses++
			break
		case 1:
			wins++
			break
		case 0:
			ties++
		}
	}

	result := &OddsResult{
		A:     a,
		B:     b,
		Total: total,
		Wins:  float64(wins) / float64(total),
		Ties:  float64(ties) / float64(total),
		Loses: float64(loses) / float64(total),
	}

	nodeHTTP.RespondJSON(resp, result)
}

func (nodeHTTP *NodeHTTP) RandomHand(resp http.ResponseWriter, req *http.Request) {
	dealer := model.NewDealer()
	board := dealer.Share(5)

	h := make([]pocketHand, 9)
	i := 0
	for i < 9 {
		pocket := dealer.Deal(2)
		cards := append(*pocket, *board...)
		//log.Printf("dealer=%s", dealer.String())
		hand, _ := poker.Detect[ranking.High](&cards)
		h[i].Pocket = *pocket
		h[i].Hand = hand
		i++
	}
	deal := dealHand{
		Board:   *board,
		Pockets: h,
	}

	nodeHTTP.RespondJSON(resp, deal)
}

func (nodeHTTP *NodeHTTP) GenerateDeck(resp http.ResponseWriter, req *http.Request) {
	s, _ := json.Marshal(model.NewDealer())
	resp.Write([]byte(s))
}

func (nodeHTTP *NodeHTTP) Deal(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Bet(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Discard(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Muck(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Pot(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Stage(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Results(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) KnownHands(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Join(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Leave(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Rebuy(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) AddOn(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Seating(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Wait(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (nodeHTTP *NodeHTTP) Stats(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}
