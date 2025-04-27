package complexity_analysis

/* for 循环 */
func forLoop(n int) int {
	res := 0
	// 循环求和 1, 2, ..., n-1, n
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

/* while 循环 */
func whileLoop(n int) int {
	res := 0
	// 初始化条件变量
	i := 1
	// 循环求和 1, 2, ..., n-1, n
	for i <= n {
		res += i
		// 更新条件变量
		i++
	}
	return res
}

/* while 循环（两次更新） */
func whileLoopII(n int) int {
	res := 0
	// 初始化条件变量
	i := 1
	// 循环求和 1, 4, 10, ...
	for i <= n {
		res += i
		// 更新条件变量
		i++
		i *= 2
	}
	return res
}
