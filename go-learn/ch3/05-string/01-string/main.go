package main

import (
	"fmt"
	"unicode/utf8"
)

func splice() {
	s := "hello, world"
	fmt.Println(len(s))             // "12"
	fmt.Println(string(s[0]), s[7]) // "104 119" ('h' and 'w')
	fmt.Println(s[0:5])             // "hello"

	fmt.Println(s[:5]) // "hello"
	fmt.Println(s[7:]) // "world"
	fmt.Println(s[:])  // "hello, world"
}

func utf() {
	// 1. 直接使用中文字符串
	s1 := "世界"
	fmt.Println(s1) // 输出: 世界

	// 2. 使用 UTF-8 字节序列
	s2 := "\xe4\xb8\x96\xe7\x95\x8c"
	fmt.Println(s2) // 输出: 世界

	// 3. 使用 Unicode 16-bit 转义序列
	s3 := "\u4e16\u754c"
	fmt.Println(s3) // 输出: 世界

	// 4. 使用 Unicode 32-bit 转义序列
	s4 := "\U00004e16\U0000754c"
	fmt.Println(s4) // 输出: 世界
}

func whatRune() {
	// ✅ 合法的 rune 表示
	r1 := '世'                            // 直接使用字符
	r2 := '\u4e16'                       // Unicode 转义
	r3 := '\U00004e16'                   // 完整 Unicode 转义
	fmt.Printf("%c %c %c\n", r1, r2, r3) // 世 世 世

	// ✅ 合法的字符串字节序列
	str := "\xe4\xb8\x96\xe7\x95\x8c" // "世界"
	fmt.Println(str)                  // 世界

	// ❌ 无效的 rune 尝试
	// rErr := '\xe4\xb8\x96' // 编译错误

	// ✅ 单字节 rune (仅限 0-255)
	rByte := '\xe4'           // 相当于 0xE4 (228)
	fmt.Printf("%c\n", rByte) // 输出 � (无效字符)
}

func RuneBinary() {
	// 方法1：直接使用十进制整数
	r1 := rune(19990)      // '世' 的 Unicode 码点十进制值
	fmt.Printf("%c\n", r1) // 输出: 世

	// 方法2：类型转换
	r2 := rune(0x4E16)     // 十六进制 → 十进制自动转换
	fmt.Printf("%c\n", r2) // 输出: 世

	// 方法3：计算得出
	r3 := rune(20000 + 16) // 20016
	fmt.Printf("%c\n", r3) // 输出: 丰

	// 与其他表示法比较
	r4 := '世'
	r5 := '\u4e16'
	fmt.Println(r1 == r4 && r1 == r5) // 输出: true
}

func StringLen() {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:]) // utf8.DecodeRuneInString函数用于解码一个字符串的第一个字符，并返回该字符对应的rune和其长度。
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串。
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// 用一个简单的循环来统计字符串中字符的数目
	n := 0
	for _, _ = range s {
		n++
	}
	// for range s { // 可以忽略不需要的变量
	// 	n++
	// }
	fmt.Println("n = ", n)
}

func RuneAndUtf8() {
	// "program" in Japanese katakana
	s := "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Println(r)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
	fmt.Println(string("\\u30d7"))
	fmt.Println(rune('\u30d7'), string('\u30ed'), string(rune(12464)), '\u30e9', '\u30e0')
	// fmt.Println(string(12464)) // グ // 从未键入的 int 转换为字符串时，得到的是一个符文字符串，而不是一个数字字符串
	fmt.Println(string(r))
}

func main() {
	// splice()

	// Immutable()
	// Immutable_num()
	// Immutable_array()

	// utf()

	// whatRune()

	// RuneBinary()

	// fmt.Println(HasPrefix("hello", "he"))
	// fmt.Println(HasSuffix("hello", "lo"))
	// fmt.Println(Contains("hello", "ll"))

	// StringLen()
	RuneAndUtf8()
}
