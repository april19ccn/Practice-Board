package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// --- --- 数组的定义
	var a [3]int              // 定义长度为3的int型数组，元素全部为0
	var b = [...]int{1, 2, 3} // 定义长度为3的int型数组，元素为1, 2, 3
	// 下标2 定义为3， 下标1 定义为2
	var c = [...]int{2: 3, 1: 2} // 定义长度为3的int型数组，元素为0, 2, 3
	// 下标4 定义为5，之后从 下标4 开始
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组，元素为1, 2, 0, 0, 5, 6

	// 下标2 定义为3， 下标1 定义为2, 这时候在增加数据会报错：数组或切片字面量中重复索引 2
	// var e = [...]int{2: 3, 1: 2, 5, 5}    // 定义长度为3的int型数组，元素为0, 2, 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// --- --- 数组指针
	var a1 = [...]int{1, 2, 3}
	var b1 = &a1

	fmt.Println(a1[0], a1[1])
	fmt.Println(b1[0], b1[1])

	for i, v := range b1 {
		fmt.Println(i, v)
	}

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}

	// 验证数组指针是否指向第一个元素
	arr := [3]int{10, 20, 30}
	arrPtr := &arr

	fmt.Printf("arr address: %p\n", &arr)    // 数组变量的地址
	fmt.Printf("arrPtr value: %p\n", arrPtr) // 数组指针的值（指向数组开头）
	fmt.Printf("&arr[0]: %p\n", &arr[0])     // 第一个元素的地址

	fmt.Printf("arrPtr == &arr[0]: %t\n", uintptr(unsafe.Pointer(arrPtr)) == uintptr(unsafe.Pointer(&arr[0])))

	// --- --- size
	// x := [100]struct {
	// 	x, y int
	// 	t    bool
	// }{struct {
	// 	x, y int
	// 	t    bool
	// }{x: 1, y: 1000, t: true}}
	// println(unsafe.Sizeof(x))

	// --- --- print
	fmt.Printf("b: %T\n", b)  // b: [3]int
	fmt.Printf("b: %#v\n", b) // b: [3]int{1, 2, 3}
}
