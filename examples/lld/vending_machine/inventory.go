package main

import (
	"errors"
	"sync"
)

type Item struct {
	code     string
	name     string
	quantity int
	price    int
	category string
}

type Inventory struct {
	item map[string]Item
	mu   sync.RWMutex
}

func NewInventory() *Inventory {
	return &Inventory{
		item: make(map[string]Item),
	}
}

func (i *Inventory) AddItem(item Item) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.item[item.code] = item
}

func (i *Inventory) RemoveItem(code string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	delete(i.item, code)
}

func (i *Inventory) GetItem(code string) (*Item, bool) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	item, ok := i.item[code]
	return &item, ok
}

func (i *Inventory) HasStock(code string) bool {
	i.mu.RLock()
	defer i.mu.RUnlock()
	item, isExist := i.item[code]
	if !isExist {
		return false
	}
	return item.quantity > 0
}

func (i *Inventory) ReduceStock(code string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	item, isExist := i.item[code]
	if !isExist || item.quantity <= 0 {
		return errors.New("out of stock")
	}
	item.quantity--
	return nil
}

func (i *Inventory) ReStock(code string, quantity int) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	item, isExist := i.item[code]
	if !isExist {
		return errors.New("first Add Item to the list")
	}
	item.quantity += quantity
	return nil
}
