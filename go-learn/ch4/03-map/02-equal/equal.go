package main

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(equal(map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2, "c": 3}))
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
}
