package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 验证编译器对字符串的优化
// 相同字符串共享一个底层数组
// 这个优化范围有多大？
// 关于Go语言中相同字符串字面量常量共享存储的范围，
// 经过分析，其核心优化机制主要作用于模块（module）内部，跨模块的优化情况则相对复杂，取决于具体编译器实现。
func checkAddress() {
	s := "hello, world"
	hello := s[:5]
	world := s[7:]
	fmt.Println(hello)
	fmt.Println(world)

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("address(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)        // 4298661077
	fmt.Println("address(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Data)      // 4298661077
	fmt.Println("address(word):", (*reflect.StringHeader)(unsafe.Pointer(&world)).Data) // 4298661084
	fmt.Println("address(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Data)      // 4298661084

	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5
}

// 处理中文字符
func handleChinese() {
	fmt.Printf("%#v\n", []byte("hello, 世界"))

	fmt.Println("\xe4\xb8\x96") // 打印“世”
	fmt.Println("\xe7\x95\x8c") // 打印“界”
}

// 测试损坏的Unicode
func destoryUnicode() {
	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
		fmt.Println(i, c)
	}
	// 0 65533  // \uFFF，对应�
	// 1 0      // 空字符
	// 2 0      // 空字符
	// 3 30028  // 界
	// 6 97     // a
	// 7 98     // b
	// 8 99     // c
}

// 遍历原始的字节码
func forByte() {
	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}
	// 0 228
	// 1 184
	// 2 150
	// 3 231
	// 4 149
	// 5 140
	// 6 97
	// 7 98
	// 8 99

	fmt.Println("----------------")

	const s = "\xe4\x00\x00\xe7\x95\x8cabc"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %x\n", i, s[i])
	}
	// 0 e4
	// 1 0
	// 2 0
	// 3 e7
	// 4 95
	// 5 8c
	// 6 61
	// 7 62
	// 8 63
}

// Rune
func whatRune() {
	fmt.Printf("%#v\n", []rune("世界"))             // []int32{19990, 30028}
	fmt.Printf("%#v\n", string([]rune{'世', '界'})) // 世界
}

func main() {
	// checkAddress()
	// handleChinese()
	// destoryUnicode()
	// forByte()
	// whatRune()

	forOnString("Hello, 世界", func(i int, r rune) {
		fmt.Printf("位置 %d: 字符 '%c' (Unicode: %U)\n", i, r, r)
	})
}
