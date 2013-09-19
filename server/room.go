package server

import (
	"log"
	"os"
	"path"
)

import (
	"gopoker/event"
	"gopoker/model"
	"gopoker/play"
	rpc_service "gopoker/server/noderpc"
	"gopoker/storage"
)

// Room - play wrapper
type Room struct {
	Guid model.Guid
	*play.Play
	Waiting  map[string]*Session
	Watchers map[string]*Session
}

// NewRoom - create new room
func NewRoom(guid model.Guid, variation model.Variation, stake model.Stake) *Room {
	newPlay := play.NewPlay(variation, stake)

	room := &Room{
		Guid:     guid,
		Play:     newPlay,
		Waiting:  map[string]*Session{},
		Watchers: map[string]*Session{},
	}

	return room
}

func (r *Room) createLogger(dir string) {
	f, err := os.OpenFile(path.Join(dir, string(r.Guid)+".log"), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal("cant open logger file", err)
	}
	loggerObserver := event.NewObserver(play.NewLogger(f))
	r.Broadcast.Broker.Bind(event.Default, event.Subscriber{
		Role:    event.System,
		Key:     "log",
		Channel: &loggerObserver.Recv,
	})
}

func (r *Room) createStorage(ps *storage.PlayHistory) {
	storageObserver := event.NewObserver(play.NewStorage(ps))
	r.Broadcast.Broker.Bind(event.Default, event.Subscriber{
		Role:    event.System,
		Key:     "history",
		Channel: &storageObserver.Recv,
	})
}
