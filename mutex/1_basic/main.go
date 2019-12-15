package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Bank struct {
	L        sync.Mutex
	accounts map[string]float64
}

func NewBank() *Bank {
	return &Bank{accounts: make(map[string]float64)}
}

func (b *Bank) Add(user string, amount float64) {
	b.L.Lock()
	defer b.L.Unlock()
	b.accounts[user] += amount
}

func (b *Bank) ShowAccounts() {
	b.L.Lock()
	defer b.L.Unlock()
	fmt.Println(b.accounts)
}

func main() {
	rand.Seed(time.Now().Unix())
	bank := NewBank()
	wg := sync.WaitGroup{}
	start := time.Now()

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			bank.Add("Alice", 1)
			wg.Done()
		}()
	}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			bank.ShowAccounts()
			wg.Done()
		}()
	}

	wg.Wait()
	log.Printf("Took %v on make all transactions", time.Since(start))
}
