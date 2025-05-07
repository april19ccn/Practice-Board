package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// data, err := ioutil.ReadFile(filename) // ioutil.ReadFile 已被舍弃
		data, err := os.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		// strings.Split 函数把字符串分割成子串的切片
		// ReadFile 函数返回一个字节切片（byte slice），必须把它转换为 string
		for _, line := range strings.Split(string(data), "\r\n") { // Windows下测试注意换行是否为\r\n
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("字符串：%s\t重复数目：%d\n", line, n)
		}
	}
}
