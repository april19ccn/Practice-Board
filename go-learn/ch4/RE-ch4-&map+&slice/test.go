// https://golang-china.github.io/gopl-zh/ch4/ch4-03.html
// 禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。
// 那为什么可以对slice中的元素取地址，slice不是也有可能会 “随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效” 吗？

// 因此，根本原因不在于“是否会重新分配内存”，而在于重新分配内存后，旧地址是否还有效以及这种有效性是否对程序员是清晰可控的。
// Slice通过返回新值并保持旧值不变，保证了旧地址的有效性范围是清晰的。Map的透明扩容机制则导致旧地址必然失效且无法安全追踪。
// 为了语言的安全性和简单性，Go选择了允许slice元素取址而禁止map元素取址。

// 该程序是为了观察 slice 内存变化， 对某个元素的地址和内容的影响
package main

import "fmt"

func main() {
	slices := make([]string, 1, 2)
	slices[0] = "a"

	t := &slices[0]
	fmt.Println(t)  // 0xc0000583c0
	fmt.Println(*t) // a

	slicesa := []string{"e", "d", "f", "g", "h", "i", "x", "y", "z", "j", "k", "l", "m", "n", "o"}
	slices = append(slices, slicesa...)

	slices[0] = "1"
	fmt.Println(t)  // 0xc0000583c0
	fmt.Println(*t) // a

	new_t := &slices[0]
	fmt.Println(new_t)  // 0xc000070000
	fmt.Println(*new_t) // 1
}
