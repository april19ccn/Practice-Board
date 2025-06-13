// 一次遍历实现 rotate
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

	arr := make([]int, len(s))

	if direction == "right" {
		n = len(s) - n
	}

	rightLen := len(s) - n
	for index := range s {
		if index < rightLen {
			arr[index] = s[n+index]
		} else {
			arr[index] = s[index-rightLen]
		}
		// fmt.Println(arr)
	}

	return arr
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// s = rotate("left", s, 2)
	s = rotate("right", s, 2)
	fmt.Println(s)
}
