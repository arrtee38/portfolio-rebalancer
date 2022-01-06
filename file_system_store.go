package main

import (
	//"encoding/json"
	"io"
)

type FileSystemAssetStore struct {
	database io.ReadSeeker
}

func (f *FileSystemAssetStore) GetPortfolio() []Asset {
	f.database.Seek(0, 0)
	portfolio, _ := NewPortfolio(f.database)
	return portfolio
}

func (f *FileSystemAssetStore) GetAssetAmount(name string) float64 {
	var amount float64
	
	for _, asset := range f.GetPortfolio() {
		if asset.Name == name {
			amount = asset.Amount
			break
		}
	}
	
	return amount
}