package main

import "fmt"

// 在计算机科学中，原地算法（in-place algorithm，也称就地算法）是不需要额外的数据结构就能变换输入数据的算法。
// 不过，分配少量空间给部分辅助变量是被允许的。
func deleteRepeat(s []string) []string {
	res := s[:0]
	for i, v := range s {
		if i > 0 && v == s[i-1] {
			continue
		}
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(deleteRepeat([]string{"a", "a", "b", "c", "c", "c"}))       // a、b、c
	fmt.Println(deleteRepeat([]string{"a1", "a1", "b2", "c2", "c1", "c2"})) // a、b、c
}
