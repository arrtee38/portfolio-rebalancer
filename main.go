package main

import (
	"log"
	"net/http"
)

//type InMemoryAssetStore struct{}

//func (i *InMemoryAssetStore) RecordAmount(name string) {}

//func (i *InMemoryAssetStore) GetAssetAmount(name string) int {
//	return 123
//}

func main() {
	server := &AssetServer{NewInMemoryAssetStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}