package play

import (
	"gopoker/storage"
)

// Storage - storage for play data
type Storage struct {
	*storage.PlayStore
	Current *storage.Play
}

// NewStorage - create new storage
func NewStorage(ps *storage.PlayStore) *Storage {
	return &Storage{
		PlayStore: ps,
	}
}
