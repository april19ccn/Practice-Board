package main

import "fmt"

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// 旋转slice
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])  // 转化成slice
	fmt.Println(a) // "[5 4 3 2 1 0]"

	// slice元素循环向左旋转n个元素
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	fmt.Println(s) // "[1 0 2 3 4 5]"
	reverse(s[2:])
	fmt.Println(s) // "[1 0 5 4 3 2]"
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	// slice元素循环向右旋转n个元素
	k := []int{0, 1, 2, 3, 4, 5}
	reverse(k)
	reverse(k[:2])
	reverse(k[2:])
	fmt.Println(k)
}
