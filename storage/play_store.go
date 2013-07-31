package storage

import (
	"log"
)

import (
	"labix.org/v2/mgo"
)

type PlayStoreConfig struct {
	Host     string
	Database string
}

type PlayStore struct {
	Session *mgo.Session
	Config  *PlayStoreConfig
}

func OpenPlayStore(config *PlayStoreConfig) *PlayStore {
	session, err := mgo.Dial(config.Host)
	if err != nil {
		log.Fatalf("[mongo] error: %s", err)
	}

	session.SetMode(mgo.Monotonic, true)
	store := &PlayStore{
		Session: session,
		Config:  config,
	}

	return store
}

func (ps *PlayStore) Close() {
	ps.Session.Close()
}

func (ps *PlayStore) Database() *mgo.Database {
	return ps.Session.DB(ps.Config.Database)
}
