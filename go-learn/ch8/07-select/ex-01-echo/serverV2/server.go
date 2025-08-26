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

	var closeOnce sync.Once
	closeConn := func() {
		closeOnce.Do(func() {
			fmt.Println("关闭连接！")
			c.Close()
		})
	}

	restore := make(chan struct{})
	go func() {
		timer := time.NewTimer(10 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-timer.C: // 可以优化一下这个 看看重复创建有什么影响
				fmt.Println("客户端10s未发送消息，已关闭")
				closeConn()
				return
			case <-restore:
				if !timer.Stop() { // 如果计时器已过期或停止，则返回 false ,所以要清空过期值
					<-timer.C
				}
				timer.Reset(10 * time.Second)
			}
		}
	}()

	input := bufio.NewScanner(c)
	for input.Scan() {
		// 非阻塞发送重置信号
		select {
		case restore <- struct{}{}:
		default: // 如果通道已满，说明已有待处理的重置信号
		}

		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, &wg) // 并发使用回声，使其更贴近真实
	}

	// !+ Closer
	go func() {
		wg.Wait()
		closeConn()
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
