package main

import (
	"fmt"
	"unicode"
)

func deleteRepeatSpace(b []byte) []byte {
	for i, v := range b {
		if unicode.IsSpace(rune(v)) {
			fmt.Println(i)
			j := i
			for j < len(b) && unicode.IsSpace(rune(b[j])) {
				j++
			}

			if i+1 < len(b) {
				fmt.Printf("%c \n", b[i+1:])
				fmt.Printf("%c \n", b[j:])
				copy(b[i+1:], b[j:])
				b = b[:len(b)-(j-1-i)]
				fmt.Printf("%c %d\n", b, len(b))
			}
		}
	}
	return b
}

func main() {
	// s := deleteRepeatSpace([]byte{'1', '2', ' ', ' ', '3', ' ', 'e', ' ', '\xe4'})
	s := deleteRepeatSpace([]byte{'1', '2', ' ', ' ', '3', ' ', ' ', ' ', ' ', '5', ' '})
	fmt.Printf("%c %d", s, len(s))
}
