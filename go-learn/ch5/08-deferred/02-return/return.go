package main

import "fmt"

// ----------------
// 关键点：匿名函数是一个闭包，它直接捕获变量 result（不是值拷贝）。
// 执行顺序：
// 		return x + x 将 8 赋值给返回值变量 result。
// 		然后执行 defer 函数，此时 result 的值已是 8。
// 		因此输出：double(4) = 8。
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

// ----------------
// 关键点：defer test(x, result) 中的 x 和 result 会在 defer 语句执行时立即求值。
// 执行顺序：
// 		当 defer 语句执行时（函数开头处），result 还是初始值 0。
// 		test(x, result) 的参数被固定为 x=4, result=0（值拷贝）。
// 		return x + x 将 8 赋值给 result，但这对已固定的参数无影响。
// 		函数退出时执行 test(4, 0)，输出：double(4) = 0。
func test(x, result int) {
	fmt.Printf("double(%d) = %d\n", x, result)
}

func double2(x int) (result int) {
	defer test(x, result)
	return x + x
}

// ----------------
func double3(x int) (result int) {
	var double = func() { fmt.Printf("double(%d) = %d\n", x, result) }
	defer double()
	return x + x
}

func main() {
	_ = double(4)
	// Output:
	// "double(4) = 8"

	_ = double2(4)
	// Output:
	// "double(4) = 0"

	_ = double3(4)
	// Output:
	// "double(4) = 8"
}
