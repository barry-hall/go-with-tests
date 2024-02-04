package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.Store.GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	if name == "Barley" {
		return "20"
	}

	if name == "River" {
		return "10"
	}

	return "0"
}
