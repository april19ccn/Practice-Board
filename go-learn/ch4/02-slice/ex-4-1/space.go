// 答案：
// 利用copy实现删除空格
// https://www.cnblogs.com/white-album2/p/16018590.html
package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := a([]byte{'1', '2', ' ', ' ', '3', ' ', ' ', ' ', ' ', '5', ' '})
	// s := a([]byte{'1', '2', ' ', ' ', '3', ' ', 'e', ' ', '\xe4'})
	fmt.Printf("%c %d", s, len(s))
}

func a(s []byte) []byte {
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsSpace(rune(s[i])) {
			if unicode.IsSpace(rune(s[i+1])) {
				fmt.Printf("%c \n", s[i+1:])
				fmt.Printf("%c \n", s[i+2:])
				copy(s[i+1:], s[i+2:])
				fmt.Printf("%c", s)
				s = s[:len(s)-1]
				fmt.Printf("%c", s)
				fmt.Println("---------")
				i--
			}
		}
	}
	return s
}
