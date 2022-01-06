package main

import (
	"encoding/json"
	"io"
)

type FileSystemAssetStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemAssetStore) GetPortfolio() Portfolio {
	f.database.Seek(0, 0)
	portfolio, _ := NewPortfolio(f.database)
	return portfolio
}

func (f *FileSystemAssetStore) GetAssetAmount(name string) float64 {
	asset := f.GetPortfolio().Find(name)
	
	if asset != nil {
		return asset.Amount
	}
	
	return 0
}

func (f *FileSystemAssetStore) RecordAmount(name string) {
	portfolio := f.GetPortfolio()
	asset := portfolio.Find(name)
	
	if asset != nil {
		asset.Amount += 1.0
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(portfolio)
}