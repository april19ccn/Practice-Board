// 练习 7.10： sort.Interface类型也可以适用在其它地方。
// 编写一个 IsPalindrome(s sort.Interface) bool 函数表明序列s是否是回文序列，
// 换句话说反向排序不会改变这个序列。假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。
package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - 1 - i

		// Less 报告索引为 i 的元素是否必须在索引为 j 的元素之前排序。
		// 如果 Less（i， j） 和 Less（j， i） 都是假的，则索引 i 和 j 处的元素被认为是相等的。
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}

	return true
}

func main() {
	// 测试
	t1 := []int{1, 2, 3, 2, 1}
	fmt.Println(IsPalindrome(sort.IntSlice(t1))) // true

	t2 := []string{"hello", "world", "hello"}
	fmt.Println(IsPalindrome(sort.StringSlice(t2))) // true

	t3 := []string{"hello", "world", "hell"}
	fmt.Println(IsPalindrome(sort.StringSlice(t3))) // false
}
