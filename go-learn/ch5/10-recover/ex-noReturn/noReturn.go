// 练习5.19： 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。

package main

func test() (res int) {
	defer func() {
		if err := recover(); err != nil {
			res = -1
		}
	}()

	panic("z")
}

func main() {
	res := test()
	println(res)
}
