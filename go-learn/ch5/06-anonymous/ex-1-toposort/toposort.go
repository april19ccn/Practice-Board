// 练习5.11： 现在线性代数的老师把微积分设为了前置课程。完善topSort，使其能检测有向图中的环。
package main

import (
	"fmt"
)

// prereqs记录了每个课程的前置课程
var Prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for k, v := range topoSort(Prereqs) {
		fmt.Println(v, k)
	}
}

func topoSort(m map[string][]string) map[string]int {
	order := make(map[string]int)
	seen := make(map[string]bool)

	index := 1
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[item] = index
				index++
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	fmt.Println(keys)

	visitAll(keys)
	return order
}
