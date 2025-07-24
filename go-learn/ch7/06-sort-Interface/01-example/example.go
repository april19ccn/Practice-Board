package main

import (
	"fmt"
	"sort"
)

// sort.Interface 接口：
// package sort

// type Interface interface {
//     Len() int
//     Less(i, j int) bool // i, j are indices of sequence elements
//     Swap(i, j int)
// }

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	strs := StringSlice{"c", "a", "b", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	sort.Sort(StringSlice(strs))
	fmt.Println(strs)

	strs1 := StringSlice{"c", "a", "b", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	sort.Strings(strs1)
	fmt.Println(strs1)
}
