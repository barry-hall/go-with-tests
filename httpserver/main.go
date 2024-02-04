package main

import (
	. "go-with-tests/httpserver/server"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	http.ListenAndServe(":5000", handler)
}
