package main

import (
	srv "go-with-tests/httpserver/server"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &srv.PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))
}
