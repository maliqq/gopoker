package storage

import (
	"time"
)

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

import (
	"gopoker/protocol/message"
)

type PlayStoreConfig struct {
	Host     string
	Database string
}

type PlayStore struct {
	Session *mgo.Session
	Config  *PlayStoreConfig
}

type Play struct {
	Id    bson.ObjectId `bson:"_id"`
	Start time.Time
	Stop  time.Time
	*message.Play
	Log []*message.Message
}

func NewObjectId() bson.ObjectId {
	return bson.NewObjectId()
}

func OpenPlayStore(config *PlayStoreConfig) (*PlayStore, error) {
	session, err := mgo.Dial(config.Host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	store := &PlayStore{
		Session: session,
		Config:  config,
	}

	return store, nil
}

func (ps *PlayStore) Close() {
	ps.Session.Close()
}

func (ps *PlayStore) Database() *mgo.Database {
	return ps.Session.DB(ps.Config.Database)
}
