package main

import "fmt"

func main() {
	original := []byte{'H', 'e', 'l', 'l', 'o', ',', ' ', 'W', 'o', 'r', 'l', 'd', '!'}
	slice := original[7:12] // 新切片 slice 指向 "World" (W, o, r, l, d)

	fmt.Println(string(slice)) // 输出: World

	// 通过新切片修改共享的底层数组
	slice[0] = 'w' // 将 'W' 改为 'w'

	// 查看原始切片（共享的底层数组也被修改了）
	fmt.Println(string(original)) // 输出: Hello, world!
}
