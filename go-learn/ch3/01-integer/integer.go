package main

import "fmt"

func main() {
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
