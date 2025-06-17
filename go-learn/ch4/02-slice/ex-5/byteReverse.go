package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverseByte(b []byte) {
	for i := 0; i < len(b); i++ {
		_, size := utf8.DecodeRune(b[i:]) // utf8.DecodeRune函数用于解码一个字符串的第一个字符，并返回该字符对应的rune和其长度。
		reverse(b[i : i+size])
	}
	reverse(b)
}

func main() {
	// b := []byte("hello, 世界")
	b := []byte("hello, 世界 abcdefg 123 %*0")
	fmt.Println(b) // [104 101 108 108 111 44 32 (228 184 150) (231 149 140)]
	reverseByte(b)
	fmt.Println(string(b))
}

// 参考答案：
//  练习 4.7
// func ReverseStringBytesSlice(s []byte) {
// 	// 第 1 轮反转（字符内反转）：遍历每个 UTF-8 字符，并对每个字符的字节在该字符的索引内进行反转
// 	// 无论这个字符有多少字节，都在它们局部的索引进行反转，这是为了第二次反转做准备
// 	for i := 0; i < len(s); {
// 		_, size := utf8.DecodeRune(s[i:])
// 		reverseByteSlice(s[i : i+size])
// 		i += size
// 	}
// 	// 第 2 轮反转（整体反转）：遍历整个字节切片，对每个字节的索引进行反转
// 	// 如此，在第 1 轮反转的那些“多字节的 UTF-8 字符”的字节在这一轮反转后，又变回正确的顺序
// 	reverseByteSlice(s)
// 	fmt.Println("Reversed byte slice: ", s)
// }
// // 辅助函数（for ReverseStringBytesSlice）：反转一个 UTF-8 字符的字节
// func reverseByteSlice(s []byte) {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }
