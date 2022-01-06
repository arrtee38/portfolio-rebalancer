package main

import (
	"encoding/json"
	"io"
)

type FileSystemAssetStore struct {
	database io.ReadWriteSeeker
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

func (f *FileSystemAssetStore) RecordAmount(name string) {
	portfolio := f.GetPortfolio()
	
	for i, asset := range portfolio {
		if asset.Name == name {
			portfolio[i].Amount += 1.0
		}
	}
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(portfolio)
}