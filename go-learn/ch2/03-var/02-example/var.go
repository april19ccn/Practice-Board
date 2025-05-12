package main

import "fmt"

// 「1」指针作为返回值
func f() *int {
	v := 1
	return &v
}

// 「2」指针作为参数
func incr(p *int) int {
	// *p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func main() {
	// 验证「1」
	var p = f()
	*p = 2
	fmt.Println(*p)

	// 验证「2」
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
}
