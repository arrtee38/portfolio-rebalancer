package main

import (
	"log"
	"net/http"
)

type InMemoryAssetAmount struct{}

func (i *InMemoryAssetAmount) GetAssetAmount(name string) int {
	return 123
}

func main() {
	server := &AssetServer{&InMemoryAssetAmount{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}