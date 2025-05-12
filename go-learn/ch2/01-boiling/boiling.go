// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212.0 // 常量boilingF是在包一级范围声明语句声明的

func main() {
	// f和c两个变量是在main函数内部声明的声明语句声明的
	// 局部声明的名字就只能在函数内部很小的范围被访问
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

	// Output:
	// boiling point = 212°F or 100°C
}
