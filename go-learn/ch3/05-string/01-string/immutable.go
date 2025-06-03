// 探究字符串底层原理
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 注意区分变量地址和数据地址

// 获取字符串底层数据指针
func stringPtr(s string) uintptr {
	return (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
}

func Immutable() {
	s := "left foot"
	fmt.Println("s 开始前地址：", &s)
	// ptr1 := stringPtr(s) // 获取变量s中存的字符串的数据地址
	ptr1 := unsafe.Pointer(&s) // 获取变量s的地址
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

// 对应基本类型 int bool 等，数据地址就是变量地址
func Immutable_num() {
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

func Immutable_array() {
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

// Immutable_array 可以看到数组两个底层数据地址是不一致的，可以推出 t := s，是 s 的拷贝值

// Immutable 可以看到底层的字符串地址是一样的，而 t := s, 并不是 s 的拷贝值，而是因为底层字符串不可以改变的性质，保证了 t 依然是旧的字符串
