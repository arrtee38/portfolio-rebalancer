package main

import (
	"testing"
	"io"
	"io/ioutil"
	"os"
)

func TestFileSystemStore(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[
            {"Name": "Stonks", "Amount": 30.0},
            {"Name": "Cryptos", "Amount": 70.0}]`)
	defer cleanDatabase()

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
		want := 70.0
		assertAmountEquals(t, got, want)
	})

	t.Run("store amounts for existing assets", func(t *testing.T) {
		store.RecordAmount("Cryptos")

		got := store.GetAssetAmount("Cryptos")
		want := 71.0
		assertAmountEquals(t, got, want)
	})
}

func assertAmountEquals(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}