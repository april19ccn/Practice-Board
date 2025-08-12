package main

import "fmt"

// 在主协程里向有缓存的通道存和取
// ✔
// func main() {
// 	ch := make(chan int, 1)
// 	ch <- 1
// 	// go func() { ch <- 1 }()
// 	//close(ch)
// 	<-ch
// 	close(ch)
// 	fmt.Println("End")
// }

// 在协程里向有缓存的通道存，在主协程里取
// ✔
// func main() {
// 	ch := make(chan int, 1)
// 	// ch <- 1
// 	go func() { ch <- 1 }()
// 	//close(ch)
// 	<-ch
// 	close(ch)
// 	fmt.Println("End")
// }

// 在协程里向有缓存的通道存，但主协程立刻关闭通道
// panic: send on closed channel
// func main() {
// 	ch := make(chan int, 1)
// 	// ch <- 1
// 	go func() { ch <- 1 }()
// 	close(ch)
// 	fmt.Println(<-ch) // 信道关闭 返回零值，不在阻塞
// 	// close(ch)
// 	fmt.Println("End")
// }

// ❌ 主协程直接接收通道
// all goroutines are asleep
func main() {
	ch := make(chan int, 1) // 有没有缓冲区都一样
	fmt.Println(<-ch)
	go func() { ch <- 1 }()
	close(ch)
	fmt.Println("End")
}

// ---------------------------------------------------------

// panic: send on closed channel
// func main() {
// 	ch := make(chan int, 1)
// 	ch <- 2
// 	go func() { ch <- 3 }()
// 	<-ch
// 	// fmt.Println(<-ch)
// 	// fmt.Println(<-ch)
// 	// time.Sleep(1 * time.Second) // 同理 加上时间就好
// 	close(ch)
// 	fmt.Println("End")
// }

// 对比第一个为什么这个不报错 	fmt.Println(<-ch) 可以提供极其短暂的时间窗口 允许协程运行，ch <- 3
// 同理，在第一个例子里 time.Sleep(1 * time.Second) 增加这一句在 close(ch) 之前，提供时间窗口
// ✔
// 2
// End
// func main() {
// 	ch := make(chan int, 1)
// 	ch <- 2
// 	go func() { ch <- 3 }()
// 	fmt.Println(<-ch)
// 	// fmt.Println(<-ch)
// 	close(ch)
// 	fmt.Println("End")
// }

// ---------------------------------------------------------
// 无缓存
// fatal error: all goroutines are asleep - deadlock!
// 无缓冲通道必须用于goroutine间的同步通信，发送和接收操作必须发生在不同的goroutine中，否则会导致死锁。
// func main() {
// 	ch := make(chan int)
// 	ch <- 1 // 当执行ch <- 1时，由于没有其他goroutine在接收数据，main goroutine会永久阻塞在这里。
// 	// go func() { ch <- 1 }()
// 	// close(ch)
// 	<-ch
// 	close(ch)
// 	fmt.Println("End")
// }
