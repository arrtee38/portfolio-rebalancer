package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Portfolio []Asset

func (p Portfolio) Find(name string) *Asset {
	for i, a := range p {
		if a.Name == name {
			return &p[i]
		}
	}
	return nil	
}

func NewPortfolio(rdr io.Reader) ([]Asset, error) {
	var portfolio []Asset
	err := json.NewDecoder(rdr).Decode(&portfolio)
	if err != nil {
		err = fmt.Errorf("problem parsing portfolio, %v", err)
	}

	return portfolio, err
}