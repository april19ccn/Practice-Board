package main

import (
	"example/learn/ch9/01-race-conditions/01-bank/bank"
	"fmt"
	"time"
)

func main() {
	// Alice:
	go func() {
		bank.Deposit(200)                // A1
		fmt.Println("=", bank.Balance()) // A2
	}()

	// Bob:
	go bank.Deposit(100) // B

	// Yu:
	go func() {
		time.Sleep(3 * time.Second)
		bank.Deposit(200)
	}()

	select {}
}
