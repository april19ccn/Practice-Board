package main

import (
	"fmt"
)

func test() (int, string) {
	return 123, "123"
}

func main() {
	var m = make(map[string]int)

	m["hello"] = 1
	m["world"] = 2

	key := "yu"

	v := m[key] // map查找，失败时返回零值
	fmt.Println(v)

	// _, ok := m[key] // map返回2个值
	_, ok := m[""], false // map返回1个值
	// _ = m[""]            // map返回1个值
	fmt.Println(ok)

	// 一个函数调用出现在元组赋值右边的表达式中时（译注：右边不能再有其它表达式），左边变量的数目必须和右边一致
	// f, err, f1, err1 := test(), test()
}
