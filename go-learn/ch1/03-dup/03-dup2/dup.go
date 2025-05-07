// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 测试命令：
// go run dup.go test.txt

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			/**
			os.Open 函数返回两个值:
			第一个值是被打开的文件（*os.File），其后被 Scanner 读取。
			第二个值是内置 error 类型的值。
			- 如果 err 等于内置值nil（译注：相当于其它语言里的 NULL），那么文件被成功打开。
			-- 读取文件，直到文件结束，然后调用 Close 关闭该文件，并释放占用的所有资源
			- 如果 err 不等于 nil，那么文件打开失败
			*/
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)

			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("字符串：%s\t重复数目：%d\n", line, n)
		}
	}
}

// 注意 countLines 函数在其声明前被调用。
// 函数和包级别的变量（package-level entities）可以任意顺序声明，并不影响其被调用。（译注：最好还是遵循一定的规范）
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() { // 对于 test.txt 也是一行一行读取
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
