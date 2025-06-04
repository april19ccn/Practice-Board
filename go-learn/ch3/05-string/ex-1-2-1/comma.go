// 将1-2的返回改为string，测试性能差异（验证string的性能消耗）
package main

import (
	"bytes"
	"strings"
)

// “12345”处理后成为“12,345”
func comma(s string) string {
	symbol := handlSymbol(s)
	comma := handleComma(s)
	dot := handleDot(s)

	buf := bytes.Buffer{}
	buf.WriteString(symbol)
	buf.WriteString(comma)
	buf.WriteString(dot)

	return buf.String()
}

func handlSymbol(s string) string {
	if s[0] == '-' || s[0] == '+' {
		return string(s[0])
	}
	return ""
}

func handleComma(s string) string {
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	var buf bytes.Buffer
	n := strings.LastIndex(s, ".")
	if n < 0 {
		n = len(s)
	}

	// 处理前面不足3位的部分
	remainder := n % 3
	if remainder == 0 {
		remainder = 3
	}

	buf.WriteString(s[:remainder])

	// 处理剩余部分，每3位加一个逗号
	for i := remainder; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

func handleDot(s string) string {
	dotData := bytes.Buffer{}
	dotIndex := strings.LastIndex(s, ".")

	if dotIndex >= 0 {
		dotData.WriteString(s[dotIndex:])
	}

	return dotData.String()
}
