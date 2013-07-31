package poker

import (
	"encoding/csv"
	"io"
	"os"
	"testing"
)

import (
	"gopoker/poker/hand"
)

func TestHandPrintString(t *testing.T) {
	high, _ := ParseCards("8dTh")
	value, _ := ParseCards("2d5s7hJdQd")
	h := Hand{
		High:  high,
		Value: value,
	}

	for rank, _ := range hand.Ranks {
		h.Rank = rank
		t.Logf(h.PrintString())
	}
}

func testHand(t *testing.T, record []string) {
	t.Logf("record=%#v", record)

	cards, _ := ParseCards(record[0])
	ranking := hand.Ranking(record[1])
	rank := hand.Rank(record[2])
	value, _ := ParseCards(record[3])
	high, _ := ParseCards(record[4])
	kicker, _ := ParseCards(record[5])

	hand, err := Detect[ranking](&cards)
	if err != nil {
		t.Fatalf("Detect error: %s", err)
	}
	t.Logf("hand=%s", hand)

	if hand.Rank != rank {
		t.Fatalf("rank mismatch: %s vs %s", hand.Rank, rank)
	}

	if !hand.Value.Equal(value) {
		t.Fatalf("value mismatch: %s vs %s", hand.Value, value)
	}

	if !hand.High.Equal(high) {
		t.Fatalf("rank mismatch: %s vs %s", hand.High, high)
	}

	if !hand.Kicker.Equal(kicker) {
		t.Fatalf("rank mismatch: %s vs %s", hand.Kicker, kicker)
	}
}

func TestHandData(t *testing.T) {
	file, err := os.Open("hand_test.csv")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	reader.Comment = '#'

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatalf("Error: %s", err)
		}
		testHand(t, record)
	}
}
