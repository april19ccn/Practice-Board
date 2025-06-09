package main

import "fmt"

func main() {
	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f - 32))     // "0";   5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float
}
