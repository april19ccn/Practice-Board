package main

import (
	"reflect"
	"unicode/utf8"
	"unsafe"
)

// for range对字符串的遍历模拟
func forOnString(s string, forBody func(i int, r rune)) {
	for i := 0; len(s) > 0; {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}

// string -> bytes
// [​]byte(s)转换模拟
func str2bytes(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		p[i] = c
	}
	return p
}

// bytes -> string
// string(bytes)转换模拟
func bytes2str(s []byte) (p string) {
	data := make([]byte, len(s))

	// 原书代码：用的loop，应该是为了突出将数据逐一复制到字符串中
	// for i, c := range s {
	// 	data[i] = c
	// }
	copy(data, s)

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	// uintptr 是一种足够大的整数类型，可以存储任何指针的位模式。
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)

	return p
}

// string -> runes
// [​]rune(s)转换模拟
func str2runes(s string) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	return []rune(p)
}

// runes -> string
// string(runes)转换模拟
func runes2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}
