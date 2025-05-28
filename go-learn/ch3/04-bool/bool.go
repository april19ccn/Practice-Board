package main

import "fmt"

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool { return i != 0 }

func main() {
	var x = false
	fmt.Println(btoi(x))

	var y = 1
	fmt.Println(itob(y))

	// fmt.Println(int(x) == 1) // cannot convert x (variable of type bool) to type int
}
