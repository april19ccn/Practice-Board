// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 内置函数 make 创建空 map, map 存储了键/值（key/value）的集合
	counts := make(map[string]int)

	// 程序使用短变量声明创建 bufio.Scanner 类型的变量 input
	// os 包是 Go 语言标准库中的一个系统相关包，os.Stdin 代表标准输入，也就是程序运行时从键盘输入的数据。
	// 它是一个 *os.File 类型的对象，可用来读取输入数据。
	input := bufio.NewScanner(os.Stdin)

	/**
	Scan 函数在读到一行时返回 true，不再有输入时返回 false。
	发送 EOF 信号终止输入：
	- Linux/Mac:Ctrl + D
	- Windows:Ctrl + Z
	*/
	for input.Scan() {
		// input.Text() 返回当前行的文本
		// map 中不含某个键时不用担心，首次读到新行时，等号右边的表达式 counts[line] 的值将被计算为其类型的零值，对于 int 即 0。
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			// fmt.Printf 函数对一些表达式产生格式化输出
			// %d          整数
			// %s          字符串
			fmt.Printf("字符串：%s\t重复数目：%d\n", line, n)
		}
	}
}
