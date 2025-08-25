// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 223.

// Reverb1 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	// !+
	defer wg.Done()
	// !+
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup

	restore := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.After(5 * time.Second): // 可以优化一下这个 看看重复创建有什么影响
				fmt.Println("客户端5s未发送消息，已关闭")
				c.Close()
				return
			case <-restore:
				continue
			}
		}
	}()

	input := bufio.NewScanner(c)
	for input.Scan() {
		restore <- struct{}{}
		// !+
		wg.Add(1)
		// !+
		go echo(c, input.Text(), 1*time.Second, &wg) // 并发使用回声，使其更贴近真实
	}

	// !+ Closer
	go func() {
		wg.Wait()
		c.Close()
	}()
	// !+

	// !- 因为client端一旦终止客户端输入，关闭写入通道会结束 input.Scan() 的阻塞，导致在协程还在运行时，就关闭了连接。
	// c.Close()
	// !-
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
