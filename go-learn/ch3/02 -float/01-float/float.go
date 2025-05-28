package main

import (
	"fmt"
	"math"
)

func main() {
	// 常量表达式比较
	fmt.Println(0.1+0.2 == 0.3) // 输出：true（高精度计算）

	// 变量运算比较
	a, b := 0.1, 0.2
	fmt.Println(a+b == 0.3) // 输出：false（双精度浮点误差）

	var f float32 = 16777216 // 1 << 24
	fmt.Println(f)
	fmt.Println(f + 1)
	fmt.Println(f == f+1) // "true"!

	// 验证无穷大是否唯一
	var z float64
	fmt.Println(math.Inf(1) == 1/z)
	fmt.Println(math.Inf(1) == math.Inf(1))  // true（同符号无穷相等）
	fmt.Println(math.Inf(1) == math.Inf(-1)) // false（符号不同）

	// 检查正无穷
	x := math.Inf(1)               // 正无穷
	fmt.Println(math.IsInf(x, 12)) // true
}
