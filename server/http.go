package server

import (
	"encoding/json"
	"gopoker/play"
	"gopoker/poker"
	"net/http"
)

type HttpService struct {
	node *Node
}

type CompareResult struct {
	A      *poker.Hand
	B      *poker.Hand
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

func (service HttpService) Hello(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) DetectHand(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	r := q.Get("ranking")
	if r == "" {
		r = "high"
	}

	ranking := poker.High
	switch ranking {
	case "badugi":
		ranking = poker.Badugi
	}

	if c := q.Get("cards"); c != "" {
		cards, err := poker.ParseCards(c)
		if err != nil {
			return resp.Write([]byte(err.Error()))
		}

		hand, err2 := ranking.Detect(cards)
		if err2 != nil {
			return resp.Write([]byte(err2.Error()))
		}
		s, _ := json.Marshal(hand)

		resp.Write([]byte(s))
	} else {
		resp.Write([]byte("no cards specified"))
	}
}

func (service HttpService) CompareHands(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	a, _ := poker.ParseCards(q.Get("a"))
	b, _ := poker.ParseCards(q.Get("b"))

	dealer := play.NewDealer()
	dealer.Burn(a)
	dealer.Burn(b)
	board := dealer.Share(5)
	c1 := append(*a, *board...)
	c2 := append(*b, *board...)
	h1, _ := poker.High.Detect(&c1)
	h2, _ := poker.High.Detect(&c2)

	s, _ := json.Marshal(&CompareResult{
		h1, h2, h1.Compare(h2),
	})

	resp.Write([]byte(s))
}

func (service HttpService) CalculateOdds(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	a, _ := poker.ParseCards(q.Get("a"))
	b, _ := poker.ParseCards(q.Get("b"))

	total := 10000
	wins, ties, loses := 0, 0, 0
	for i := 0; i <= total; i++ {
		dealer := play.NewDealer()
		dealer.Burn(a)
		dealer.Burn(b)
		board := dealer.Share(5)
		c1 := append(*a, *board...)
		c2 := append(*b, *board...)
		h1, _ := poker.High.Detect(&c1)
		h2, _ := poker.High.Detect(&c2)

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
	s, _ := json.Marshal(&OddsResult{
		a, b, total,
		float64(wins) / float64(total),
		float64(ties) / float64(total),
		float64(loses) / float64(total),
	})

	resp.Write([]byte(s))
}

func (service HttpService) RandomHand(resp http.ResponseWriter, req *http.Request) {
	dealer := play.NewDealer()
	board := dealer.Share(5)

	h := make([]pocketHand, 9)
	i := 0
	for i < 9 {
		pocket := dealer.Deal(2)
		cards := append(*pocket, *board...)
		//log.Printf("dealer=%s", dealer.String())
		hand, _ := poker.High.Detect(&cards)
		h[i].Pocket = *pocket
		h[i].Hand = hand
		i++
	}
	deal := dealHand{
		Board:   *board,
		Pockets: h,
	}

	s, _ := json.Marshal(deal)

	resp.Write([]byte(s))
}

func (service HttpService) GenerateDeck(resp http.ResponseWriter, req *http.Request) {
	s, _ := json.Marshal(play.NewDealer())
	resp.Write([]byte(s))
}

func (service HttpService) GetDeal(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) Bet(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) Discard(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) Muck(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) GetPot(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) GetStage(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) GetResults(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) GetKnownHands(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) GetTable(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) JoinTable(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) LeaveTable(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) Rebuy(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) AddOn(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) GetTableSeating(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) Wait(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}

func (service HttpService) GetTableStats(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello, world!"))
}
