// 验证如果那两个无缓存channel同时处于发送状态，select怎么处理
package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go func() {
		fmt.Println("channel-1 准备就绪")
		channel1 <- "测试11111"
	}()
	go func() {
		fmt.Println("channel-2 准备就绪")
		channel2 <- "测试22222"
	}()

	time.Sleep(2 * time.Second)

	for {
		select {
		case s1 := <-channel1:
			fmt.Println(s1)
		case s2 := <-channel2:
			fmt.Println(s2)
		}
	}
}
