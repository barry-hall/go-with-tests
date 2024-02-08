package server

import "sync"

// NewInMemoryPlayerStore initialises an empty player store.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct {
	Store map[string]int
	// A mutex is used to synchronize read/write access to the map
	lock sync.RWMutex
}

// RecordWin will record a player's win.
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.Store[name]++
}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.Store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player

	for name, wins := range i.Store {
		league = append(league, Player{name, wins})
	}
	return league
}
