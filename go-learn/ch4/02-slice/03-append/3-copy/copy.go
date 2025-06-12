package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5} // 底层数组
	src := arr[1:4]             // [2,3,4] (len=3)
	dst := arr[0:3]             // [1,2,3] (len=3)

	fmt.Println(src)
	fmt.Println(dst)

	fmt.Println("结果：")
	fmt.Println(copy(dst, src))
	fmt.Println(arr)
	fmt.Println(dst)
	fmt.Println(src)
	// 结果：
	// arr = [2, 3, 4, 4, 5]
	// dst = [2,3,4] (覆盖前3位)
	// src = [3,4,4] (共享数组被修改)
}
