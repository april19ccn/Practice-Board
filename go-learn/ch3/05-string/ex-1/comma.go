// 存在大量string的创建操作
package main

import (
	"bytes"
	"slices"
	"strings"
)

// “12345”处理后成为“12,345”
func comma(s string) string {
	s, symbol := handlSymbol(s)

	s1, dotData := handleDot(s)

	s2 := handleComma(s1)

	buf := bytes.Buffer{}
	buf.WriteString(symbol)
	buf.WriteString(s2)

	if dotData != "" {
		buf.WriteString(".")
		buf.WriteString(dotData)
	}

	return buf.String()
}

func handlSymbol(s string) (string, string) {
	if s[0] == '-' || s[0] == '+' {
		return s[1:], s[0:1]
	}
	return s, ""
}

func handleComma(s string) string {
	var buf bytes.Buffer
	var arr []string

	for i := len(s) - 3; i >= 0; i -= 3 {
		arr = append(arr, s[i:i+3])
	}

	if (len(s) % 3) != 0 {
		arr = append(arr, s[:len(s)%3])
	}

	slices.Reverse(arr)
	// fmt.Println(arr)

	for k, v := range arr {
		buf.WriteString(v)
		if k != len(arr)-1 {
			buf.WriteByte(',')
		}
	}

	return buf.String()
}

func handleDot(s string) (string, string) {
	dotData := ""
	dotIndex := strings.LastIndex(s, ".")

	if dotIndex >= 0 {
		dotData = s[dotIndex+1:]
		s = s[:dotIndex]
	}

	return s, dotData
}
