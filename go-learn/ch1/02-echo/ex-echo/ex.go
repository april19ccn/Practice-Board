package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func printCommandName() {
	fmt.Println("命令名：")
	fmt.Println(os.Args[0])
}

func printCommandArgs() {
	fmt.Println("命令行参数：")
	for index, arg := range os.Args[1:] {
		fmt.Println("第" + strconv.Itoa(index) + "个参数: " + arg)
	}
}

func testGreateStringTime() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Printf("耗时：%d μs\n", time.Since(start).Microseconds()) // time.Since(start): 计算从 start 时间点到当前时间的时间差（即耗时）

	start1 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("耗时：%d μs\n", time.Since(start1).Microseconds())
}

// 练习1.3 通过 基准测试 验证时间效率
func connectStringJoin(str []string) string {
	return strings.Join(str[0:], " ")
}

func connectStringFor(str []string) string {
	s, sep := "", ""
	for _, arg := range str[0:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))

	// 练习 1.1
	printCommandName()

	// 练习 1.2
	printCommandArgs()

	// 练习 1.3 通过 time 对比时间效率
	testGreateStringTime()
}
