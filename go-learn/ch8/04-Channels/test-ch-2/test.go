// https://github.com/gopl-zh/gopl-zh.github.com/discussions/168#discussioncomment-9973515
// https://github.com/lazybearlee

// 操作不同状态的chan会引发三种行为：

// panic
// 	1. 向已经关闭的通道写数据
// 	2. 重复关闭通道
// 阻塞
// 	1. 向未初始化的通道写/读数据
// 	2. 向缓冲区已满的通道写入数据
// 	3. 通道中没有数据，读取该通道
// 非阻塞
// 	1. 读取已经关闭的通道，这个操作会返回通道元素类型的零值（可用comma, ok语法）
//	2. 向有缓冲且没有满的通道读/写

package main

import (
	"testing"
)

func TestChanOperateStatus(t *testing.T) {
	t.Run("向已经关闭的通道写数据", func(t *testing.T) {
		ch := make(chan int)
		close(ch) // 关闭通道
		ch <- 1   // 这里会引发panic，因为向已关闭的通道发送数据
		// panic: send on closed channel [recovered]
	})
	t.Run("重复关闭通道", func(t *testing.T) {
		ch := make(chan int)
		close(ch) // 第一次关闭通道
		close(ch) // 再次关闭通道会引发panic
		// panic: close of closed channel [recovered]
	})
	t.Run("向未初始化的通道写/读数据", func(t *testing.T) {
		var ch chan int
		go func() {
			ch <- 1
			// x := <-ch
		}()
		_ = <-ch
		// fatal error: all goroutines are asleep - deadlock!
	})
	t.Run("向缓冲区已满的通道写入数据", func(t *testing.T) {
		ch := make(chan int, 1)
		ch <- 1 // 第一次写入，缓冲区未满
		ch <- 2 // 这里会阻塞，因为缓冲区已满，没有goroutine读取数据
		// fatal error: all goroutines are asleep - deadlock!
	})
	t.Run("通道中没有数据，读取该通道", func(t *testing.T) {
		ch := make(chan int)
		_ = <-ch // 这里会阻塞，因为没有goroutine发送数据到通道
		// fatal error: all goroutines are asleep - deadlock!
	})
	t.Run("读取已经关闭的通道，这个操作会返回通道元素类型的零值（可用comma, ok语法）", func(t *testing.T) {
		ch := make(chan int)
		close(ch)     // 关闭通道
		x, ok := <-ch // x 将会是int类型的零值，ok 将会是false
		expectx, expectok := 0, false
		if ok != expectok && x != expectx {
			t.Errorf("expect 0, false, get %d, %t\n", x, ok)
		}
	})
	t.Run("向有缓冲且没有满的通道写，向有缓冲且不为空的通道读", func(t *testing.T) {
		ch := make(chan int, 2) // 1 也不会堵塞
		ch <- 1                 // 写入数据，不会阻塞
		_ = <-ch                // 读取数据，不会阻塞
	})
}
