package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}

// // 如果buffer改成2 select会随机执行
// func main() {
// 	ch := make(chan int, 2)
// 	for i := 0; i < 10; i++ {
// 		select {
// 		case x := <-ch:
// 			fmt.Println(x) // "0" "2" "4" "6" "8"
// 		case ch <- i:
// 		}
// 	}
// }
