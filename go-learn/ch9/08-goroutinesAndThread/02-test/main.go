package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// 创建两个无缓冲 channel
	ping := make(chan struct{})
	pong := make(chan struct{})

	var count int
	var wg sync.WaitGroup
	wg.Add(2)

	// 使用 context 来控制 goroutine 的退出
	ctx, cancel := context.WithCancel(context.Background())

	// 启动第一个 goroutine (ping)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ping:
				// 收到 ping，发送 pong
				select {
				case pong <- struct{}{}:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	// 启动第二个 goroutine (pong)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-pong:
				// 收到 pong，计数并发送 ping
				count++
				select {
				case ping <- struct{}{}:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	// 开始计时并启动通信
	start := time.Now()
	ping <- struct{}{} // 发送第一个 ping

	// 运行一秒钟
	time.Sleep(time.Second)

	// 安全地停止 goroutine
	cancel()
	wg.Wait()

	elapsed := time.Since(start)

	fmt.Printf("每秒通信次数: %d\n", count)
	fmt.Printf("总通信次数: %d\n", count)
	fmt.Printf("总耗时: %v\n", elapsed)
	if count > 0 {
		fmt.Printf("平均每次通信耗时: %v\n", time.Duration(int(elapsed)/count))
	}
}
