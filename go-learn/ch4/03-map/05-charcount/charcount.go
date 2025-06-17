// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters Unicode字符的计数
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings UTF-8 编码的长度计数
	invalid := 0                    // count of invalid UTF-8 characters 无效 UTF-8 字符计数

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error // 如果编码的 rune 无效，则消耗 1 个字节并返回 unicode。ReplacementChar （U+FFFD）
		if err == io.EOF {         // EOF 是 "读取 "在没有更多输入时返回的错误。
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 { // 代表无效代码点
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
