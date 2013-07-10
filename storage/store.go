package storage

import (
	"database/sql"
)

type Store struct {
	*sql.DB
}

const (
	DefaultDriver = "postgres"
)

func OpenStore(driver string, dataSource string) (*Store, error) {
	store = &Store{}
	store.DB, err = sql.Open(driver, dataSource)
	return store, err
}

func (store *Store) Close() {
	store.DB.Close()
}
