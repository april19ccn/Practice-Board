package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
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

// FuzzIsPalindrome 对 IsPalindrome 函数进行模糊测试
func FuzzIsPalindrome(f *testing.F) {
	// 添加种子测试用例
	seedCases := []string{
		"",                               // 空字符串
		"a",                              // 单字符
		"aa",                             // 简单回文
		"aba",                            // 奇数长度回文
		"a man a plan a canal panama",    // 带空格的回文
		"hello",                          // 非回文
		"racecar",                        // 回文
		"A man, a plan, a canal: Panama", // 带标点的回文
		"12321",                          // 数字回文（但会被过滤掉数字）
		"😀a😀",                            // 包含 Unicode 表情符号
	}

	for _, tc := range seedCases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// 调用被测试函数
		result := IsPalindrome(input)

		// 验证结果的一致性
		// 方法1：如果返回 true，验证预处理后的字符串确实是回文
		if result {
			// 手动实现预处理逻辑
			var letters []rune
			for _, r := range input {
				if unicode.IsLetter(r) {
					letters = append(letters, unicode.ToLower(r))
				}
			}

			// 验证预处理后的字符串确实是回文
			n := len(letters)
			for i := 0; i < n/2; i++ {
				if letters[i] != letters[n-1-i] {
					t.Errorf("IsPalindrome returned true for non-palindrome. Input: %q, Preprocessed: %q",
						input, string(letters))
				}
			}
		}

		// 方法2：对于某些特定情况，我们可以确定结果应该是什么
		// 例如：空字符串、单字符字符串应该总是回文
		if input == "" {
			if !result {
				t.Errorf("Empty string should always be palindrome, got false")
			}
			return
		}

		// 单字符（字母）应该总是回文
		if len(input) == 1 {
			r := rune(input[0])
			if unicode.IsLetter(r) && !result {
				t.Errorf("Single letter %q should be palindrome, got false", input)
			}
		}

		// 方法3：验证函数不会 panic
		// 模糊测试会自动捕获 panic，但我们也可以显式检查
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsPalindrome panicked with input %q: %v", input, r)
			}
		}()

		// 再次调用以确保一致性（可选）
		result2 := IsPalindrome(input)
		if result != result2 {
			t.Errorf("IsPalindrome not consistent for input %q: first call %t, second call %t",
				input, result, result2)
		}
	})
}

// 运行模糊测试
// go test -fuzz=FuzzIsPalindrome -fuzztime=30s
