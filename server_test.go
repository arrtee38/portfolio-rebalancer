package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubAssetAmount struct {
	amounts map[string]int
	amountCalls []string
}

func (a *StubAssetAmount) GetAssetAmount(name string) int {
	amount := a.amounts[name]
	return amount
}

func (a *StubAssetAmount) RecordAmount(name string) {
	a.amountCalls = append(a.amountCalls, name)
}

func TestGETAssets(t *testing.T) {
	store := StubAssetAmount{
		map[string]int{
			"Stonks": 20,
			"Cryptos":  10,
		},
		nil,
	}
	server := NewAssetServer(&store)
	
	t.Run("returns amount of Stonks", func(t *testing.T) {
		request := newGetAmountRequest("Stonks")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns amount of Crypto", func(t *testing.T) {
		request := newGetAmountRequest("Cryptos")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})


	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetAmountRequest("Bonds")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreAmounts(t *testing.T) {
	store := StubAssetAmount{
		map[string]int{},
		nil,
	}
	server := NewAssetServer(&store)

	t.Run("it records wins on POST", func(t *testing.T) {
		asset := "Stonks"

		request := newPostAmountRequest(asset)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.amountCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.amountCalls), 1)
		}

		if store.amountCalls[0] != asset {	
			t.Errorf("did not store correct winner got %q want %q", store.amountCalls[0], asset)
		}
	})
}

func TestPortfolio (t *testing.T) {
	store := StubAssetAmount{}
	server := NewAssetServer(&store)

	t.Run("it returns 200 on /portfolio", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/portfolio", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
		if got != want {
			t.Errorf("did not get correct status, got %d, want %d", got, want)
		}
}

func newGetAmountRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/assets/%s", name), nil)
	return req
}

func newPostAmountRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/assets/%s", name), nil)
	return req
}
