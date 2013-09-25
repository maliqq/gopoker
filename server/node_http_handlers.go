package server

import (
	"encoding/json"
	"net/http"
)

import (
	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

import (
	"gopoker/model"
	"gopoker/poker"
	"gopoker/poker/hand"
	"gopoker/poker/math"
	"gopoker/server/node_response"
)

// Rooms - GET /rooms
func (nodeHTTP *NodeHTTP) Rooms(resp http.ResponseWriter, req *http.Request) {
	rooms := nodeHTTP.Node.Rooms

	nodeHTTP.RespondJSON(resp, rooms)
}

// Room - GET /room/{id}
func (nodeHTTP *NodeHTTP) Room(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	room := nodeHTTP.Node.Room(model.Guid(id))

	nodeHTTP.RespondJSON(resp, room)
}

// DetectHand - GET /hand/detect?ranking=...&cards=...
func (nodeHTTP *NodeHTTP) DetectHand(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	r := q.Get("ranking")
	if r == "" {
		r = "High"
	}
	ranking := hand.Ranking(r)

	if c := q.Get("cards"); c != "" {
		cards, err := poker.ParseCards(c)
		if err != nil {
			resp.Write([]byte(err.Error()))
			return
		}

		hand, err := poker.Detect[ranking](&cards)
		if err != nil {
			resp.Write([]byte(err.Error()))
			return
		}

		nodeHTTP.RespondJSON(resp, hand)
	} else {
		resp.Write([]byte("no cards specified"))
	}
}

// CompareHands - GET /hand/compare?a=...&b=...
func (nodeHTTP *NodeHTTP) CompareHands(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	a, _ := poker.ParseCards(q.Get("a"))
	b, _ := poker.ParseCards(q.Get("b"))

	dealer := model.NewDealer()
	dealer.Burn(a)
	dealer.Burn(b)
	board := dealer.Share(5)
	c1 := append(a, board...)
	c2 := append(b, board...)
	h1, _ := poker.Detect[hand.High](&c1)
	h2, _ := poker.Detect[hand.High](&c2)

	s, _ := json.Marshal(&node_response.CompareResult{
		A:      h1,
		B:      h2,
		Board:  board,
		Result: h1.Compare(h2),
	})

	resp.Write([]byte(s))
}

// CalculateOdds - GET /hand/odds?a=...&b=...
func (nodeHTTP *NodeHTTP) CalculateOdds(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	a, _ := poker.ParseCards(q.Get("a"))
	b, _ := poker.ParseCards(q.Get("b"))

	total := 10000
	chances := math.ChancesAgainstOne{SamplesNum: total}.Preflop(a, b)

	result := &node_response.OddsResult{
		A:     a,
		B:     b,
		Total: total,
		Wins:  chances.Wins(),
		Ties:  chances.Ties(),
		Loses: chances.Loses(),
	}

	nodeHTTP.RespondJSON(resp, result)
}

// RandomHand - GET /hand/random
func (nodeHTTP *NodeHTTP) RandomHand(resp http.ResponseWriter, req *http.Request) {
	dealer := model.NewDealer()
	board := dealer.Share(5)

	h := make([]node_response.PocketHand, 9)
	i := 0
	for i < 9 {
		pocket := dealer.Deal(2)
		cards := append(pocket, board...)
		hand, _ := poker.Detect[hand.High](&cards)
		h[i].Pocket = pocket
		h[i].Hand = hand
		i++
	}
	deal := node_response.DealHand{
		Board:   board,
		Pockets: h,
	}

	nodeHTTP.RespondJSON(resp, deal)
}

// GenerateDeck - GET /deck/generate
func (nodeHTTP *NodeHTTP) GenerateDeck(resp http.ResponseWriter, req *http.Request) {
	s, _ := json.Marshal(model.NewDealer())
	resp.Write([]byte(s))
}

// Play - GET /play/{id}
func (nodeHTTP *NodeHTTP) Play(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	play, err := nodeHTTP.Node.PlayHistory.Find(id)

	if err != nil {
		glog.Errorf("[error] %s", err)

		nodeHTTP.RespondJSONError(resp, err)
	} else {
		nodeHTTP.RespondJSON(resp, play)
	}
}

// Winners - GET /play/{id}/winners
func (nodeHTTP *NodeHTTP) Winners(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	play, err := nodeHTTP.Node.PlayHistory.Find(id)

	if err != nil {
		glog.Errorf("[error] %s", err)

		nodeHTTP.RespondJSONError(resp, err)
	} else {
		nodeHTTP.RespondJSON(resp, play.Winners)
	}
}

// KnownCards - GET /play/id/known_cards
func (nodeHTTP *NodeHTTP) KnownCards(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	play, err := nodeHTTP.Node.PlayHistory.Find(id)

	if err != nil {
		glog.Errorf("[error] %s", err)

		nodeHTTP.RespondJSONError(resp, err)
	} else {
		nodeHTTP.RespondJSON(resp, play.KnownCards)
	}
}

// Bet - bet via HTTP
func (nodeHTTP *NodeHTTP) Bet(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

// Discard - discard via HTTP
func (nodeHTTP *NodeHTTP) Discard(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

// Muck - muck via HTTP
func (nodeHTTP *NodeHTTP) Muck(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

// Pot - get play pot
func (nodeHTTP *NodeHTTP) Pot(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

// Stage - get play stage
func (nodeHTTP *NodeHTTP) Stage(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

// Join - join table via HTTP
func (nodeHTTP *NodeHTTP) Join(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

// Leave - leave table via HTTP
func (nodeHTTP *NodeHTTP) Leave(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

// Rebuy - rebuy to table via HTTP
func (nodeHTTP *NodeHTTP) Rebuy(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

// AddOn - addon to table via HTTP
func (nodeHTTP *NodeHTTP) AddOn(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

// Seating - get table seating
func (nodeHTTP *NodeHTTP) Seating(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

// Wait - join waiting queue
func (nodeHTTP *NodeHTTP) Wait(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

// Stats - get table stats
func (nodeHTTP *NodeHTTP) Stats(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}
