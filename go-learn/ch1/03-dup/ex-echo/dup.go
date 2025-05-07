// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 测试命令：
// go run .\dup.go test.txt test2.txt test3.txt

func main() {
	fileStack := make(map[string]map[string]int)
	counts := make(map[string]int)
	files := os.Args[1:]

	fmt.Println(files)

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			fmt.Println(arg)

			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)

			fileStack[arg] = counts

			counts = make(map[string]int)

			f.Close()
		}
	}
	for filename, counts := range fileStack {
		for line, n := range counts {
			if n > 1 {
				fmt.Println(filename + ":")
				fmt.Printf("字符串：%s\t重复数目：%d\n", line, n)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
