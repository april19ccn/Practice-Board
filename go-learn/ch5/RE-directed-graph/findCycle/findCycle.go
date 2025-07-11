// 找到所有包含环的路径
package main

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
var Prereqs = map[string][]string{
	"algorithms": {"data structures"},

	// ----环
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	// -----

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

// sample 环 A-B—C
var graph1 = map[string][]string{
	"K": {"A", "B"},
	"A": {"B", "D"},
	"B": {"C"},
	"C": {"A"},
}

// sample 无环
var graph2 = map[string][]string{
	"K": {"A"},
	"A": {"B"},
	"B": {"C"},
	"C": {"D"},
}

func main() {
	fmt.Println(findCycle(Prereqs))
	fmt.Println(findCycle(graph1))
	fmt.Println(findCycle(graph2))
}

func findCycle(graph map[string][]string) [][]string {
	visited := make(map[string]bool)
	cycle := make([]string, 0)
	res := make([][]string, 0)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				cycle = append(cycle, item)
				// fmt.Println(item, graph[item], cycle)
				visitAll(graph[item])
			} else {
				cycle = append(cycle, item)
				temp := make([]string, len(cycle)) // 要存slice的深拷贝，否则会受到底层数组的影响
				copy(temp, cycle)
				res = append(res, temp)
				// cycle = append(cycle, item)
			}

			// fmt.Println("------------------")
			visited[item] = false
			if len(cycle) > 0 {
				cycle = cycle[:len(cycle)-1]
			}
		}
	}

	var keys []string
	for key := range graph {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Printf("keys: %v \n", keys)

	visitAll(keys) // 第一次循环用的keys，是为了防止环外有遗漏的节点

	return res
}
