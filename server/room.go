package server

import (
	"log"
	"os"
	"path"
)

import (
	"gopoker/model"
	"gopoker/play"
	"gopoker/protocol"
	rpc_service "gopoker/server/noderpc"
	"gopoker/storage"
)

// Room - play wrapper
type Room struct {
	ID string
	*play.Play
	Waiting  map[string]*Session
	Watchers map[string]*Session
}

// NewRoom - create new room
func NewRoom(createRoom *rpc_service.CreateRoom) *Room {
	var variation model.Variation

	if createRoom.Mix != nil {
		variation = createRoom.Mix
	} else {
		variation = createRoom.Game
	}

	stake := model.NewStake(createRoom.BetSize)
	newPlay := play.NewPlay(variation, stake)

	room := &Room{
		ID:       createRoom.ID,
		Play:     newPlay,
		Waiting:  map[string]*Session{},
		Watchers: map[string]*Session{},
	}

	return room
}

func (r *Room) createLogger(dir string) {
	f, err := os.OpenFile(path.Join(dir, r.ID+".log"), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal("cant open logger file", err)
	}
	logger := play.NewLogger(f)
	r.Broadcast.Broker.Bind(protocol.Observer, "logger", &logger.Recv)
}

func (r *Room) createStorage(ps *storage.PlayStore) {
	storage := play.NewStorage(ps)
	r.Broadcast.Broker.Bind(protocol.Observer, "storage", &storage.Recv)
}
