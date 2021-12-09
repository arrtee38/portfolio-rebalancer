package main

import (
	"fmt"
	"strings"
	"net/http"
)

type AssetAmount interface{
	GetAssetAmount(name string) int
}

type AssetServer struct {
	store AssetAmount
}

func (a *AssetServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	asset := strings.TrimPrefix(r.URL.Path, "/assets/")
	fmt.Fprint(w, a.store.GetAssetAmount(asset))
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
