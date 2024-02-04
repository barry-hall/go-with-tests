package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Barley": 20,
			"River":  10,
		},
	}
	server := &PlayerServer{&store}
	t.Run("returns Barleys score", func(t *testing.T) {
		request := newGetScoreRequest("Barley")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertBodyResponse(t, response.Body.String(), "20")

	})
	t.Run("returns Rivers score", func(t *testing.T) {
		request := newGetScoreRequest("River")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertBodyResponse(t, response.Body.String(), "10")

	})

}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertBodyResponse(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
