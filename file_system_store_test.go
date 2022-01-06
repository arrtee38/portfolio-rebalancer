package main

import (
	"testing"
	"strings"
)

func TestFileSystemStore(t *testing.T) {
	database := strings.NewReader(`[
            {"Name": "Stonks", "Amount": 30.0},
            {"Name": "Cryptos", "Amount": 70.0}]`)
	store := FileSystemAssetStore{database}
		
	t.Run("portfolio from a reader", func(t *testing.T) {
		got := store.GetPortfolio()

		want := []Asset{
			{"Stonks", 30.0},
			{"Cryptos", 70.0},
		}

		assertPortfolio(t, got, want)

		//check that re-read starts at beginning of file
		got = store.GetPortfolio()
		assertPortfolio(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		got := store.GetAssetAmount("Cryptos")
		assertAssetAmount(t, got, 70.0)
	})
}

func assertAssetAmount(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}