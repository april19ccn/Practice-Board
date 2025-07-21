// 练习 7.1： 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type Counter struct {
	words int
	lines int
}

func (c *Counter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		fmt.Println(word)
		c.words++
	}

	// 统计行数：换行符数量 + 结尾非换行符的额外一行
	// c.lines = bytes.Count(p, []byte{'\n'})
	// if len(p) > 0 && p[len(p)-1] != '\n' {
	// 	c.lines++
	// }

	pstr := string(p)
	parr := strings.Split(pstr, "\n")
	c.lines += len(parr)
	fmt.Println(parr)
	if c.lines > 0 && parr[len(parr)-1] == "" {
		c.lines--
	}

	return len(p), nil
}

func main() {
	var c Counter
	c.Write([]byte("hello, world\ntest1\ntest2\n\n123\n\n"))
	fmt.Println(c)

	c = Counter{} // reset the counter
	var test = "hello, world\ntest1\ntest2\n\n123\n\n"
	fmt.Fprintf(&c, "Add+1, %s", test)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
