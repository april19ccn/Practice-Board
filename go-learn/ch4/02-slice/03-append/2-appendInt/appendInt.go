package main

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) { // 每次调用appendInt函数，必须先检测slice底层数组是否有足够的容量来保存新添加的元素
		// 如果有足够空间的话，直接扩展slice（依然在原有的底层数组之上）
		z = x[:zlen]
	} else {
		// 如果没有足够的增长空间的话，appendInt函数则会先分配一个足够大的slice用于保存新的结果，先将输入的x复制到新的空间
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)

		// 内置的copy函数可以方便地将一个slice复制另一个相同类型的slice。
		copy(z, x)
	}

	z[len(x)] = y
	return z
}

func main() {
	appendInt([]int{1, 2, 3}, 4)
}
