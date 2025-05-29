package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func splice() {
	s := "hello, world"
	fmt.Println(len(s))             // "12"
	fmt.Println(string(s[0]), s[7]) // "104 119" ('h' and 'w')
	fmt.Println(s[0:5])             // "hello"

	fmt.Println(s[:5]) // "hello"
	fmt.Println(s[7:]) // "world"
	fmt.Println(s[:])  // "hello, world"
}

// 获取字符串底层数据指针
func stringPtr(s string) uintptr {
	return (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
}

func immutable() {
	s := "left foot"
	fmt.Println("s 开始前地址：", &s)
	// ptr1 := stringPtr(s)
	ptr1 := unsafe.Pointer(&s)
	fmt.Println("ptr1 开始前地址：", ptr1)

	t := s
	fmt.Println("t 开始前地址：", &t)
	// ptr2 := stringPtr(t)
	ptr2 := unsafe.Pointer(&t)
	fmt.Println("ptr2 开始前地址：", ptr2)

	x := &s
	fmt.Println("x 开始前地址：", &x)
	ptr3 := unsafe.Pointer(&x)
	fmt.Println("ptr3 开始前地址：", ptr3)

	s += ", right foot"

	fmt.Println(s) // "left foot, right foot"
	fmt.Println("s 开始后地址：", &s)

	fmt.Println(t) // "left foot"
	fmt.Println("t 开始后地址：", &t)

	fmt.Println(*x)
	fmt.Println("x 开始后地址：", &x)
}

func immutable_num() {
	s := 1234
	fmt.Println("s 开始前地址：", &s)

	t := s
	fmt.Println("t 开始前地址：", &t)

	s++

	fmt.Println(s) // "1235"
	fmt.Println("s 开始后地址：", &s)

	fmt.Println(t) // "1234"
	fmt.Println("t 开始后地址：", &t)
}

func immutable_array() {
	s := [3]int{1, 2, 3}
	fmt.Println("s 开始前地址：", &s[0])
	// 获取数组首元素地址（即底层数据起始指针）
	ptrS := unsafe.Pointer(&s[0])
	fmt.Printf("s 底层指针: %p\n", ptrS)

	t := s
	fmt.Println("t 开始前地址：", &t)
	ptrT := unsafe.Pointer(&t[0])
	fmt.Printf("t 底层指针: %p\n", ptrT) // 不同的地址！

	s[0]++
	fmt.Println(s) // "[2 2 3]"
	fmt.Println("s 开始后地址：", &s)
	fmt.Println(t) // "[1 2 3]"
	fmt.Println("t 开始后地址：", &t)
}

func main() {
	splice()
	immutable()
	immutable_num()
	immutable_array()
}
