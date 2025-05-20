package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 2.3
func PopCountLoop(x uint64) int {
	result := 0
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

// 2.4
func PopCountFor(x uint64) int {
	result := 0
	for x != 0 {
		result += int(byte(x & 1))
		x >>= 1
	}
	return result
}

// 2.5
func PopCountDel(x uint64) int {
	result := 0
	for x != 0 {
		x = x & (x - 1)
		result++
	}
	return result
}

func main() {
	fmt.Println(PopCount(184467440737095))
	fmt.Println(PopCountLoop(184467440737095))
	fmt.Println(PopCountFor(184467440737095))
	fmt.Println(PopCountDel(184467440737095))
}
