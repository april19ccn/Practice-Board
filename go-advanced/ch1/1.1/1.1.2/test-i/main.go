package main

import "fmt"

func closureExample() func() {
	count := 0 // 外部变量

	// 闭包函数捕获了外部变量count的引用
	increment := func() {
		count++ // 直接修改外部变量
		fmt.Printf("Count: %d\n", count)
	}

	return increment
}

func testi() {
	var funcs []func()

	for i := 0; i < 3; i++ {
		// 闭包捕获的是外部变量i的引用，不是复制值！
		funcs = append(funcs, func() {
			fmt.Printf("i = %d\n", i)
		})
	}

	// 运行所有闭包函数
	for _, f := range funcs {
		f()
	}
	// 输出：
	// i = 3
	// i = 3
	// i = 3
	// 因为所有闭包都引用同一个i，循环结束时i=3

	// 不对：实际输出是：
	// i = 0
	// i = 1
	// i = 2

	// https://go.dev/doc/go1.22
	// go 1.22
	// Go 1.22 makes two changes to “for” loops.
	// - In Go 1.22, each iteration of the loop creates new variables, to avoid accidental sharing bugs.
	// - “For” loops may now range over integers.
}

func main() {
	counter := closureExample()

	counter() // 输出: Count: 1
	counter() // 输出: Count: 2
	counter() // 输出: Count: 3

	// 每次调用都修改的是同一个count变量！

	testi()
}
