package main

import (
	"testing"
	"strings"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("portfolio from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
            {"Name": "Stonks", "Amount": 30.0},
            {"Name": "Cryptos", "Amount": 70.0}]`)

		store := FileSystemAssetStore{database}

		got := store.GetPortfolio()

		want := []Asset{
			{"Stonks", 30.0},
			{"Cryptos", 70.0},
		}

		assertPortfolio(t, got, want)

		got = store.GetPortfolio()
		assertPortfolio(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Stonks", "Amount": 70.0},
			{"Name": "Cryptos", "Amount": 30.0}]`)

		store := FileSystemAssetStore{database}

		got := store.GetAssetAmount("Cryptos")

		want := 30.0

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}