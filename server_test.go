package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubAssetAmount struct {
	amounts map[string]int
}

func (a *StubAssetAmount) GetAssetAmount(name string) int {
	amount := a.amounts[name]
	return amount
}

func TestGETAssets(t *testing.T) {
	amount := StubAssetAmount{
		map[string]int{
			"Stonks": 20,
			"Cryptos":  10,
		},
	}
	server := &AssetServer{&amount}
	
	t.Run("returns amount of Stonks", func(t *testing.T) {
		request := newGetAmountRequest("Stonks")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns amount of Crypto", func(t *testing.T) {
		request := newGetAmountRequest("Cryptos")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		
		assertResponseBody(t, response.Body.String(), "10")
	})
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func newGetAmountRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/assets/%s", name), nil)
	return req
}

