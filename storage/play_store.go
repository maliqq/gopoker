package storage

import (
	"log"
	"os"
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
	Play  *message.Play `bson:"play"`
	Log   []*message.Message
}

func NewObjectId() bson.ObjectId {
	return bson.NewObjectId()
}

var Debug = false

func OpenPlayStore(config *PlayStoreConfig) (*PlayStore, error) {
	session, err := mgo.Dial(config.Host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)
	if Debug {
		mgo.SetLogger(log.New(os.Stderr, "", log.LstdFlags))
		mgo.SetDebug(true)
	}

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

func (ps *PlayStore) Collection(collection string) *mgo.Collection {
	return ps.Database().C(collection)
}

func (ps *PlayStore) FindPlayById(id string) (*Play, error) {
	var play Play

	plays := ps.Collection("plays")
	query := plays.Find(bson.M{"_id": bson.ObjectIdHex(id)})
	err := query.One(&play)

	return &play, err
}
