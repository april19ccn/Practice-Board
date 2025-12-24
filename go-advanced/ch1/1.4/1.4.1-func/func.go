package main

import "fmt"

// ------------------- 函数类型
// 具名函数
func Add1(a, b int) int {
	return a + b
}

// 匿名函数
var Add2 = func(a, b int) int {
	return a + b
}

// ------------------- 参数和返回值
// 多个参数和多个返回值
func Swap(a, b int) (int, int) {
	return b, a
}

// 可变数量的参数
// more 对应的 []int 切片类型
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

// ------------------- 空接口类型参数
func interfaceExample() {
	var a = []interface{}{123, "abc"}

	Print(a...) // 123 abc
	// => Print(123, "abc")
	Print(a) // [123 abc]
	// => Print([]interface{}{123, "abc"})

	fmt.Println(a...) // 123 abc
	fmt.Println(a)    // [123 abc]
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

// ------------------- 返回值命名
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

// 利用 defer 语句在 return 语句之后修改返回值：
func Inc() (v int) {
	defer func() { v++ }()
	return 42
}

// ------------------- for 循环的引用
func forExample() {
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
}

// 输出（版本高于Go 1.22）：
// 2
// 1
// 0

// 输出（版本低于Go 1.22）：
// 3
// 3
// 3

func main() {
	// interfaceExample()
	// fmt.Println(Inc())
	forExample()
}
