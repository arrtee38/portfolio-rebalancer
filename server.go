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
	router := http.NewServeMux()
	router.Handle("/portfolio", http.HandlerFunc(a.portfolioHandler))
	router.Handle("/assets/", http.HandlerFunc(a.assetHandler))

	router.ServeHTTP(w, r)
}

func (a *AssetServer) portfolioHandler(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
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

	if amount == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, amount)
}

func (a *AssetServer) processAmount(w http.ResponseWriter, asset string) {
	a.store.RecordAmount(asset)
	w.WriteHeader(http.StatusAccepted)
}
