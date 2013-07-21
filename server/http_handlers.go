package server

import (
	"encoding/json"
	"net/http"
)

import (
	"github.com/gorilla/mux"
)

import (
	"gopoker/calc"
	"gopoker/model"
	"gopoker/poker"
	"gopoker/poker/ranking"
	"gopoker/server/http_service"
)

func (nodeHTTP *NodeHTTP) Rooms(resp http.ResponseWriter, req *http.Request) {
	rooms := nodeHTTP.Node.Rooms

	nodeHTTP.RespondJSON(resp, rooms)
}

func (nodeHTTP *NodeHTTP) Room(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["room"]
	room := nodeHTTP.Node.Room(id)

	nodeHTTP.RespondJSON(resp, room)
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
	h1, _ := poker.Detect[ranking.High](&c1)
	h2, _ := poker.Detect[ranking.High](&c2)

	s, _ := json.Marshal(&http_service.CompareResult{
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
	chances := calc.ChancesAgainstOne{total}.Preflop(a, b)

	result := &http_service.OddsResult{
		A:     a,
		B:     b,
		Total: total,
		Wins:  float64(chances.Wins) / float64(total),
		Ties:  float64(chances.Ties) / float64(total),
		Loses: float64(chances.Loses) / float64(total),
	}

	nodeHTTP.RespondJSON(resp, result)
}

func (nodeHTTP *NodeHTTP) RandomHand(resp http.ResponseWriter, req *http.Request) {
	dealer := model.NewDealer()
	board := dealer.Share(5)

	h := make([]http_service.PocketHand, 9)
	i := 0
	for i < 9 {
		pocket := dealer.Deal(2)
		cards := append(pocket, board...)
		//log.Printf("dealer=%s", dealer.String())
		hand, _ := poker.Detect[ranking.High](&cards)
		h[i].Pocket = pocket
		h[i].Hand = hand
		i++
	}
	deal := http_service.DealHand{
		Board:   board,
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
