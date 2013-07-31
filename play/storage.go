package play

import (
	"gopoker/protocol"
	"gopoker/protocol/message"
	"gopoker/storage"
)

type Storage struct {
	*storage.PlayStore
	Recv protocol.MessageChannel
}

func NewStorage(ps *storage.PlayStore) *Storage {
	storage := &Storage{
		PlayStore: ps,
		Recv:      make(protocol.MessageChannel),
	}

	return storage
}

func (this *Storage) receive() {
	for {
		msg := <-this.Recv
		this.handle(msg)
	}
}

func (this *Storage) handle(msg *message.Message) {
	switch msg.Payload().(type) {
	case *message.PlayStart:
	}
}
