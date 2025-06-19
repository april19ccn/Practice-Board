// 更深入了解 数字、字母等等如何精确区分，比如字符0~1，字母A~Z，中文字符等等
// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters Unicode字符的计数
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings UTF-8 编码的长度计数
	invalid := 0                    // count of invalid UTF-8 characters 无效 UTF-8 字符计数

	letters := make(map[rune]int)
	han := make(map[rune]int)
	numbers := [10]int{}
	english := make(map[rune]int)

	// bufio.NewReader 这是一个带缓冲的读取器，主要用于逐行或逐字节读取数据。 注意和 bufio.NewScanner 的区别
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

		if unicode.IsLetter(r) { // 判断是否是字母（不只是英文字母）
			letters[r]++
		}
		if unicode.IsNumber(r) { // 判断是否是数字字符（不只是阿拉伯数字）
			n, _ := strconv.Atoi(string(rune(r)))
			numbers[n]++
		}
		if unicode.Is(unicode.Han, r) {
			han[r]++
		}
		if r >= 65 && r <= 122 {
			english[r]++
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

	fmt.Println("字母统计：")
	for c, n := range letters {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Println("数字统计：")
	for c, n := range numbers {
		fmt.Printf("%d\t%d\n", c, n)
	}

	fmt.Println("汉字统计：")
	for c, n := range han {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Println("英文统计：")
	for c, n := range english {
		fmt.Printf("%q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

// 测试数据
// 中文字符：今天天气怎么样？123，123
// ABCdef の の &%…… …………^&*@
