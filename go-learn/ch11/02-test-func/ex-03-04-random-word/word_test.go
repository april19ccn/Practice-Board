package word

import (
	"math/rand"
	"testing"
	"time"
)

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24 // rng.Intn(25) 会生成一个 [0, 25) 范围内的随机整数，即 0 到 24。这个值 n 将是最终回文字符串的长度。
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		// rng.Intn(0x1000) 生成一个 [0, 4096) 范围内的整数，
		// 覆盖了 Unicode 基本多文种平面的很大一部分，
		// 包括拉丁字母、西里尔字母、部分亚洲文字等。
		// 这避免了生成控制字符或更复杂的 Unicode 字符，简化了生成逻辑。
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r

		x := rng.Intn(100)
		if x < 50 {
			runes[i] = ' '
			runes[n-1-i] = ' '
		}

		z := rng.Intn(100)
		if z < 50 {
			runes[i] = '?'
			runes[n-1-i] = '?'
		}
	}

	return string(runes)
}

// 测试非回文字符串
// func randomNotPalindrome(rng *rand.Rand) string {
// 	n := rng.Intn(25)
// 	if n < 2 { // 去掉0，防止越界panic；去掉1，1个字符是回文
// 		return "Not"
// 	}
// 	runes := make([]rune, n)
// 	for i := range n {
// 		runes[i] = rune(rng.Intn(0x1000)) // random rune up to '\u0999'
// 	}
// 	// for runes[0] == runes[n-1] { // 当 n=0 时会崩溃，使用这个需要去掉0；当 n=1 时会无限循环
// 	// 	runes[0] = rune(rng.Intn(0x1000))
// 	// }
// 	return "n" + string(runes) + "t"
// }

// IsPalindrome 会去掉非字母，导致生成具有回文的字符串！
func randomNotPalindrome(rng *rand.Rand) string {
	return "n" + randomPalindrome(rng) + "t"
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		// fmt.Println(p)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}

		n := randomNotPalindrome(rng)
		if IsPalindrome(n) {
			t.Errorf("IsPalindrome(%q) = true", n)
		}
	}
}
