// 只检测是否存在环，发现立刻返回
package main

import (
	"fmt"
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

// sample 无环
var graph3 = map[string][]string{
	"K": {"A", "B"},
	"A": {"B"},
}

func main() {
	fmt.Println(checkCycle(Prereqs))
	fmt.Println(checkCycle(graph1))
	fmt.Println(checkCycle(graph2))
	fmt.Println(checkCycle(graph3))
}

func checkCycle(graph map[string][]string) bool {
	visited := make(map[string]bool)

	var visitAll func(items []string) bool
	visitAll = func(items []string) bool {
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				if visitAll(graph[item]) {
					return true
				}
			} else {
				return true
			}

			// fmt.Println("------------------")
			visited[item] = false
		}
		return false
	}

	var keys []string
	for key := range graph {
		keys = append(keys, key)
	}
	// sort.Strings(keys)
	fmt.Println(keys)

	return visitAll(keys) // 第一次循环用的keys，是为了防止环外有遗漏的节点
}
