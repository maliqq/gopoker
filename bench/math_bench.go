package main

import (
	"time"
	"fmt"
	"gopoker/model"
	"gopoker/poker/math"
)

const (
	oppNum = 6
	nano = 1000000000.0
)

func measure(samplesNum int) {
	n := 100

	var total, min, max int64
	total = 0
	min = -1
	max = 0

	for i := 0; i < n; i++ {
		dealer := model.NewDealer()
		cards := dealer.Deal(2)
		board := dealer.Share(3)
		start := time.Now().UnixNano()

		math.ChancesAgainstN{OpponentsNum: oppNum, SamplesNum: samplesNum}.WithBoard(cards, board)
		t := time.Now().UnixNano() - start
		if (min == -1) { min = t }
		if (t > max) { max = t }
		if (t < min) { min = t }
		total += t
	}

	fmt.Printf("samplesNum=%d\tn=%d\ttotal=%.3fs\tavg=%.3fs\tmin=%.3fs\tmax=%.3fs\n", 
		samplesNum, n,
		float64(total) / nano,
		float64(total) / 10 / nano,
		float64(min) / nano,
		float64(max) / nano)
}

func main() {
	measure(1)
	measure(10)
	measure(100)
	measure(1000)
	measure(10000)
	measure(100000)
}
