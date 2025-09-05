package main

import (
	"example/learn/ch9/01-race-conditions/ex-bank/bank"
	"fmt"
)

func main() {
	go func() {
		fmt.Println("11111")
		bank.Deposit(500)
		if bank.Withdraw(600) {
			fmt.Println(bank.Balance())
		}
	}()

	go func() {
		fmt.Println("22222")
		bank.Deposit(500)
		fmt.Println(bank.Balance())
	}()

	go func() {
		fmt.Println("33333")
		if bank.Withdraw(700) {
			fmt.Println(bank.Balance())
		}
	}()

	for {

	}
}
