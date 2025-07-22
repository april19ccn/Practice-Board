// 验证defer的执行顺序
// 等return执行完 执行defer
package main

import "fmt"

func test1() bool {
	fmt.Println("test1")
	return false
}

func test2() bool {
	fmt.Println("test2")
	return true
}

func deferTest() bool {
	defer test1()
	return test2()
}

func main() {
	deferTest()
}
