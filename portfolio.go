package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewPortfolio(rdr io.Reader) ([]Asset, error) {
	var portfolio []Asset
	err := json.NewDecoder(rdr).Decode(&portfolio)
	if err != nil {
		err = fmt.Errorf("problem parsing portfolio, %v", err)
	}

	return portfolio, err
}