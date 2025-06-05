package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 123
	fmt.Printf("%s\n", fmt.Sprintf("%d", i))
	fmt.Println(strconv.Itoa(i))
	fmt.Println(string(rune(i)))

	fmt.Println(strconv.FormatInt(int64(i), 2))
	fmt.Println(strconv.FormatUint(uint64(i), 2))
}
