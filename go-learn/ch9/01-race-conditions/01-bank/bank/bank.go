// Package bank implements a bank with only one account.
package bank

var balance int

func Deposit(amount int) { balance = balance + amount } // 存款
func Balance() int       { return balance }             // 查询
