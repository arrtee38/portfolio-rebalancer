package main

import (
	//"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETAssets(t *testing.T) {
	t.Run("returns amount of Stonks", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/assets/Stonks", nil)
		response := httptest.NewRecorder()

		AssetServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}