package main

import "time"

func main() {
	var x []int
	go func() { x = make([]int, 10) }()
	go func() { x = make([]int, 100) }()
	time.Sleep(2 * time.Second)
	x[99] = 1 // NOTE: undefined behavior; memory corruption possible!
}

// 有个问题，怎么复现一个 指针是 10长度的数据， 但 slice 的长度确实 100..... ，然后导致 报错？ 还是什么？