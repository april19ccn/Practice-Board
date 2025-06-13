// 利用 reverse 函数实现 rotate
package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/**
* direction: left or right
* s: slice
* n: rotate的位置
 */
func rotate(direction string, s []int, n int) []int {
	if n <= 0 || n >= len(s) {
		return s
	}

	if direction == "right" {
		reverse(s)
	}
	reverse(s[:n])
	reverse(s[n:])
	if direction == "left" {
		reverse(s)
	}

	return s
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// s = rotate("left", s, 2)
	s = rotate("right", s, 2)
	fmt.Println(s)
}
