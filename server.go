package main

import (
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
)

type AssetAmount interface {
	GetAssetAmount(name string) float64
	RecordAmount(name string)
	GetPortfolio() []Asset
}

type AssetServer struct {
	store AssetAmount
	http.Handler
}

type Asset struct {
	Name string
	Amount float64
}

const jsonContentType = "application/json"

func NewAssetServer(store AssetAmount) *AssetServer {
	a := new(AssetServer)

	a.store = store 

	router := http.NewServeMux()
	router.Handle("/portfolio", http.HandlerFunc(a.portfolioHandler))
	router.Handle("/assets/", http.HandlerFunc(a.assetHandler))

	a.Handler = router

	return a
}

func (a *AssetServer) portfolioHandler(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", jsonContentType)
		json.NewEncoder(w).Encode(a.store.GetPortfolio())
}

func (a *AssetServer) assetHandler(w http.ResponseWriter, r *http.Request) {
	asset := strings.TrimPrefix(r.URL.Path, "/assets/")

	switch r.Method {
		case http.MethodPost:
			a.processAmount(w, asset)
		case http.MethodGet:
			a.showAmount(w, asset)
	}
} 

func (a *AssetServer) showAmount(w http.ResponseWriter, asset string) {
	amount := a.store.GetAssetAmount(asset)

	if amount == 0.0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, amount)
}

func (a *AssetServer) processAmount(w http.ResponseWriter, asset string) {
	a.store.RecordAmount(asset)
	w.WriteHeader(http.StatusAccepted)
}
