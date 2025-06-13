// 利用 copy 和 append 函数实现 rotate
package main

import "fmt"

/**
* direction: left or right
* s: slice
* n: rotate的位置
 */
func rotate(direction string, s []int, n int) []int {
	if n < 0 || n >= len(s) {
		return s
	}

	var arr1, arr2 []int

	if direction == "right" {
		arr1 = make([]int, len(s)-n)
		arr2 = make([]int, n)
		copy(arr1, s[:len(s)-n])
		copy(arr2, s[len(s)-n:])
	}
	if direction == "left" {
		arr1 = make([]int, n)
		arr2 = make([]int, len(s)-n)
		copy(arr1, s[:n])
		copy(arr2, s[n:])
	}

	return append(arr2, arr1...)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// s = rotate("left", s, 2)
	s = rotate("right", s, 2)
	fmt.Println(s)
}
