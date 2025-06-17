package main

import "fmt"

var m = make(map[string]int)

func k(list []string) string {
	fmt.Printf("%q\n", list)
	return fmt.Sprintf("%q", list)
}

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func main() {
	fmt.Println(Count([]string{"a", "b", "c"}))
	Add([]string{"a", "b", "c"})
	fmt.Println(Count([]string{"a", "b", "c"}))
}
