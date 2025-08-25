package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("芜湖~🛫")
}

func main() {
	fmt.Println("Commencing countdown.")

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case finishTime := <-time.After(10 * time.Second):
		// Do nothing.
		fmt.Println(finishTime)
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

	launch()
}
