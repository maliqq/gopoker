package play

import (
	"log"
	"time"
)

import (
	"gopoker/exch"
	"gopoker/exch/message"
	"gopoker/storage"
)

// Storage - storage for play data
type Storage struct {
	*storage.PlayStore
	Current *storage.Play
	Recv    exch.MessageChannel
}

// NewStorage - create new storage
func NewStorage(ps *storage.PlayStore) *Storage {
	storage := &Storage{
		PlayStore: ps,
		Recv:      make(exch.MessageChannel),
	}

	go storage.receive()

	return storage
}

func (stor *Storage) receive() {
	for {
		msg := <-stor.Recv
		stor.handle(msg)
	}
}
