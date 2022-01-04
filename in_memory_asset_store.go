package main

func NewInMemoryAssetStore() *InMemoryAssetStore {
	return &InMemoryAssetStore{map[string]int{}}
}

type InMemoryAssetStore struct {
	store map[string]int
}

func (i *InMemoryAssetStore) GetPortfolio() []Asset {
	var portfolio []Asset
	for name, amount := range i.store {
		portfolio = append(portfolio, Asset{name, amount})
	}

	return portfolio
}

func (i *InMemoryAssetStore) RecordAmount(name string) {
	i.store[name]++
}

func (i *InMemoryAssetStore) GetAssetAmount(name string) int {
	return i.store[name]
}