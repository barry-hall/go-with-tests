package server_test

import (
	. "go-with-tests/httpserver"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Barleys score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Barley", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
