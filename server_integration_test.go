package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryAssetStore()
	server := NewAssetServer(store)
	asset := "Stonks"

	server.ServeHTTP(httptest.NewRecorder(), newPostAmountRequest(asset))
	server.ServeHTTP(httptest.NewRecorder(), newPostAmountRequest(asset))
	server.ServeHTTP(httptest.NewRecorder(), newPostAmountRequest(asset))

	t.Run("get amount", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetAmountRequest(asset))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get portfolio", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newPortfolioRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getPortfolioFromResponse(t, response.Body)
		want := []Asset{
			{"Stonks", 3},
		}
		assertPortfolio(t, got, want)
	})
}