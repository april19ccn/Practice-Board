package main

import (
	"fmt"
	"time"
)

func counter(c chan int) {
	i := 2
	for {
		c <- i
		i++
		time.Sleep(1 * time.Second)
	}
}

func filter(prime int, listen, send chan int) {
	var i int
	for {
		i = <-listen
		if i%prime != 0 {
			send <- i
		}
		time.Sleep(1 * time.Second)
	}
}

func sieve() chan int {
	c := make(chan int)
	go counter(c)

	prime := make(chan int)
	go func() {
		var p int
		var newc chan int
		for {
			p = <-c
			prime <- p
			newc = make(chan int)
			go filter(p, c, newc)
			c = newc
		}
	}()

	return prime
}

func main() {
	result := sieve()

	for i := range result {
		fmt.Println(i)
	}
}
