package server

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemPlayerStore struct {
	Database *json.Encoder
	League   League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	file.Seek(0, 0)
	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&Tape{file}),
		League:   league,
	}, nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.League
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.League.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.League
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.League = append(f.League, Player{name, 1})
	}
	f.Database.Encode(f.League)
}
