package main

import (
	"fmt"
	"sync"
)

type Bank struct {
	sync.RWMutex
	deposit float64
}

func (b *Bank) Deposits(amount float64) {
	b.Lock()
	defer b.Unlock()
	b.deposit += amount
}

func (b *Bank) Balances() float64 {
	b.RLock()
	defer b.RUnlock()
	return b.deposit
}

func (b *Bank) Withdraw(amount float64) bool {
	b.Lock()
	defer b.Unlock()
	if b.deposit-amount < 0 {
		return false
	}
	b.deposit -= amount
	return true
}

func main() {
	account := Bank{deposit: 1000.0}

	var wg sync.WaitGroup

	// 启动多个查询操作
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("查询者 %d: 余额 %.2f\n", id, account.Balances())
		}(i)
	}

	// 启动存款操作
	wg.Add(1)
	go func() {
		defer wg.Done()
		account.Deposits(200)
	}()

	// 启动取款操作
	wg.Add(1)
	go func() {
		defer wg.Done()
		account.Withdraw(150)
	}()

	wg.Wait()
	fmt.Printf("最终余额: %.2f\n", account.Balances())
}
