// 1-2改写成fp形式
package main

import (
	"bytes"
	"strings"

	F "github.com/IBM/fp-go/function"
)

var (
	Symbol = F.Curry2(handlSymbol)
	Comma  = F.Curry2(handleComma)
	Dot    = F.Curry2(handleDot)
)

// “12345”处理后成为“12,345”
func comma(s string) string {
	var buf bytes.Buffer

	result := F.Flow3(
		Symbol(s),
		Comma(s),
		Dot(s),
	)(buf)

	return result.String()
}

func handlSymbol(s string, buf bytes.Buffer) bytes.Buffer {
	if s[0] == '-' || s[0] == '+' {
		buf.Write([]byte(s[0:1]))
	}
	return buf
}

func handleComma(s string, buf bytes.Buffer) bytes.Buffer {
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

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

	return buf
}

func handleDot(s string, buf bytes.Buffer) bytes.Buffer {
	dotIndex := strings.LastIndex(s, ".")

	if dotIndex >= 0 {
		buf.WriteString(s[dotIndex:])
	}

	return buf
}
