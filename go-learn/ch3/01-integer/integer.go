package main

import "fmt"

func bit() {
	var z = 1 << 1
	fmt.Printf("%08b\n", z)

	// uint8类型值的8个独立的bit位
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	// Printf函数的%b参数打印二进制格式的数字,%08b中08表示打印至少8个字符宽度，不足的前缀部分用0填充
	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}

func float() {
	f := 3.141 // a float64
	i := int(f)
	fmt.Println(f, i) // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"

	var i16 int16 = 1300
	fmt.Printf("%016b\n", i16)  // 0000 0101 0001 0100
	fmt.Println(i16, int8(i16)) // 保留 0001 0100 = 20
	fmt.Println(i16, int32(i16))

	f2 := 1e100
	fmt.Println(int(f2))
}

func binary() {
	// fmt 的技巧
	// 1. %之后的[1]副词告诉Printf函数再次使用第一个操作数
	// 2. %后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀
	// %#[1]o #在[1]前， 不能写成 %[1]#o
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o \n", o) // "438 666 0666"
	x1 := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x1)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
}

// 字符面值通过一对单引号直接包含对应字符
// 可以通过转义的数值来表示任意的Unicode码点对应的字符
func char() {
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
}

func main() {
	bit()

	float()

	binary()

	char()
}
