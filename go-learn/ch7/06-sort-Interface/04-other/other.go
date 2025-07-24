package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{3, 1, 4, 1}

	// IntsAreSorted 报告切片 x 是否按升序排序
	fmt.Println(sort.IntsAreSorted(values)) // "false"

	sort.Ints(values)
	fmt.Println(values)                     // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"
}
