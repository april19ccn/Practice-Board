// Package word provides utilities for word games.
package word

import (
	"unicode"
)

// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) { // 使用 unicode.IsLetter() 过滤非字母字符 ！！会去掉非字母，导致生成具有回文的字符串！
			letters = append(letters, unicode.ToLower(r)) // 大小写统一
		}
	}
	// for i := range letters {
	// 	if letters[i] != letters[len(letters)-1-i] {
	// 		return false
	// 	}
	// }
	for i := 0; i < len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
