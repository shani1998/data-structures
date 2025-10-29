package main

import (
	"fmt"
	"sync"
)

type PaymentStrategy interface {
	AddAmount(amount int) int
	DisplayBalance() int
	ProcessAmount(amount int) (int, error)
}

type CoinPayment struct {
	amount int
	sync.RWMutex
}

func NewCoinPayment() *CoinPayment {
	return &CoinPayment{}
}

var _ PaymentStrategy = &CoinPayment{}

func (c *CoinPayment) AddAmount(amount int) int {
	c.Lock()
	defer c.Unlock()
	c.amount += amount
	return c.amount
}

func (c *CoinPayment) DisplayBalance() int {
	c.RLock()
	defer c.RUnlock()
	return c.amount
}

func (c *CoinPayment) ProcessAmount(amount int) (int, error) {
	c.Lock()
	defer c.Unlock()
	if c.amount < amount {
		return 0, fmt.Errorf("amount %d < amount %d", amount, c.amount)
	}
	c.amount -= amount
	return c.amount, nil
}
