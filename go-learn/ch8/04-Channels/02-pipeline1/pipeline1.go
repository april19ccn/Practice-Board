package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			if x == 100 { // 即时终止了
				close(naturals)
				break
			}
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals // 接收方依然会无限取到 零值，导致死循环无限打印
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
