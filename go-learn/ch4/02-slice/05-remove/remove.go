package main

import "fmt"

func removeOrder(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeDisorder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(removeOrder(s, 2)) // "[5 6 8 9]"

	z := []int{5, 6, 7, 8, 9}
	fmt.Println(removeDisorder(z, 2)) // "[5 6 9 8]
}
