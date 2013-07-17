package server

import (
	"encoding/json"
	"net/http"
)

import (
	"gopoker/model"
	"gopoker/poker"
	"gopoker/poker/ranking"
)

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
