package play

import (
	"gopoker/storage"
)

// Storage - storage for play data
type Storage struct {
	History *storage.PlayHistory
	Current *storage.PlayHistoryEntry
}

// NewStorage - create new storage
func NewStorage(history *storage.PlayHistory) *Storage {
	return &Storage{
		History: history,
	}
}
