package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryAssetStore()
	server := AssetServer{store}
	asset := "Stonks"

	server.ServeHTTP(httptest.NewRecorder(), newPostAmountRequest(asset))
	server.ServeHTTP(httptest.NewRecorder(), newPostAmountRequest(asset))
	server.ServeHTTP(httptest.NewRecorder(), newPostAmountRequest(asset))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetAmountRequest(asset))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}