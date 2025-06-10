package main

import (
	"crypto/sha256"
	"fmt"
)

// c1 := sha256.Sum256([]byte("x"))
// c2 := sha256.Sum256([]byte("X"))
// fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
// // Output:
// // 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
// // 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
// // false
// // [32]uint8

// 1. sha256 会随机改变吗？
// 2. 如果随机改变应该怎么测？

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func getSHA256(x, y []byte) ([32]uint8, [32]uint8) {
	return sha256.Sum256(x), sha256.Sum256(y)
}

func diffCount(c1, c2 [32]uint8) int {
	count := 0
	for i := range c1 {
		for j := 0; j < 8; j++ {
			if c1[i]&(1<<j) != c2[i]&(1<<j) {
				count++
			}
		}
	}
	return count
}

func popCount(s1, s2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		temp := s1[i] ^ s2[i] // 按位异或运算符
		count += int(pc[temp])
	}

	return count
}

func main() {
	c1, c2 := getSHA256([]byte("x"), []byte("X"))

	fmt.Println(diffCount(c1, c2))

	fmt.Println(popCount(c1, c2))
}
