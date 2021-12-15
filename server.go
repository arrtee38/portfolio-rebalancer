package main

import (
	"fmt"
	"strings"
	"net/http"
)

type AssetAmount interface{
	GetAssetAmount(name string) int
	RecordAmount(name string)
}

type AssetServer struct {
	store AssetAmount
}

func (a *AssetServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	switch r.Method {
		case http.MethodPost:
			a.processAmount(w)
		case http.MethodGet:
			a.showAmount(w, r)
	}
}

func (a *AssetServer) showAmount(w http.ResponseWriter, r *http.Request) {
	asset := strings.TrimPrefix(r.URL.Path, "/assets/")

	amount := a.store.GetAssetAmount(asset)

	if amount == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, amount)
}

func (a *AssetServer) processAmount(w http.ResponseWriter) {
	a.store.RecordAmount("RealEstate")
	w.WriteHeader(http.StatusAccepted)
}

func GetAssetAmount(name string) string {
	if name == "Stonks" {
		return "20"
	}

	if name == "Cryptos" {
		return "10"
	}

	return ""
}
