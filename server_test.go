package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"reflect"
)

type StubAssetAmount struct {
	amounts map[string]int
	amountCalls []string
	portfolio []Asset
}

func (a *StubAssetAmount) GetAssetAmount(name string) int {
	amount := a.amounts[name]
	return amount
}

func (a *StubAssetAmount) RecordAmount(name string) {
	a.amountCalls = append(a.amountCalls, name)
}

func (a *StubAssetAmount) GetPortfolio() []Asset {
	return a.portfolio
}

func TestGETAssets(t *testing.T) {
	store := StubAssetAmount{
		map[string]int{
			"Stonks": 20,
			"Cryptos":  10,
		},
		nil,
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

	t.Run("it returns 200 on /portfolio", func(t *testing.T) {
		wantedPortfolio := []Asset{
			{"Stonks", 41},
			{"Cryptos", 52},
			{"Real Estate", 7},
		}
		
		store := StubAssetAmount{
			nil,
			nil,
			wantedPortfolio,
		}
		server := NewAssetServer(&store)
		
		request := newPortfolioRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		
		got := getPortfolioFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertPortfolio(t, got, wantedPortfolio)
		assertContentType(t, response, jsonContentType)

	})
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
		t.Helper()
		if response.Result().Header.Get("content-type") != want {
			t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
		}
}

func assertPortfolio(t testing.TB, got, want []Asset) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func getPortfolioFromResponse(t testing.TB, body io.Reader) (portfolio []Asset) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&portfolio)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Asset, '%v'", body, err)
	}
	return
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

func newPortfolioRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/portfolio", nil)
	return req
}


