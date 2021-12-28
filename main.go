package main

import (
	"log"
	"net/http"
)

func main() {
	server := &AssetServer{NewInMemoryAssetStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}