package main

import (
	"time"
	"fmt"
	"gopoker/model"
	"gopoker/poker/math"
)

const (
	oppNum = 6
)

func measure(samplesNum int) {
	var total int64
	total = 0
	dealer := model.NewDealer()
	cards := dealer.Deal(2)
	board := dealer.Share(3)
	start := time.Now().UnixNano()

	math.ChancesAgainstN{OpponentsNum: oppNum, SamplesNum: samplesNum}.WithBoard(cards, board)
	t := time.Now().UnixNano() - start
	total += t

	fmt.Printf("samplesNum=%d time=%.3fs\n", samplesNum, float64(total) / 1000000000.0)
}

func main() {
	measure(1)
	measure(10)
	measure(100)
	measure(1000)
	measure(10000)
	measure(100000)
}
