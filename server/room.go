package server

import (
	"os"
	"path"
)

import (
	"github.com/golang/glog"
)

import (
	"gopoker/event"
	"gopoker/model"
	"gopoker/play"
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
		glog.Fatalf("cant open logger file", err)
	}

	logger := play.NewLogger(f)
	logger.Subscribe(r.Hub)
}

func (r *Room) createStorage(ps *storage.PlayHistory) {
	storage := play.NewStorage(ps)
	storage.Subscribe(r.Hub)
}
