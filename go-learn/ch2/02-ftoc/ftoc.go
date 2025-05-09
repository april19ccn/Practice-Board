// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0

	// 在 Go 语言中，%g 是格式化输出中的一种占位符，主要用于格式化浮点数（float32 和 float64）。
	// 它会根据数值的大小和精度，自动选择 更简洁的表示方式，可能是 科学计数法（%e） 或 普通小数（%f）。
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
