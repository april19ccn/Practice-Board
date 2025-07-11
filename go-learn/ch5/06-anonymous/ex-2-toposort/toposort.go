// 练习5.10： 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。
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

func main() {
	if checkCycle(Prereqs) {
		fmt.Println("存在环")
		return
	}
	for k, v := range topoSort(Prereqs) {
		fmt.Println(v, k)
	}
}

// 检测是否存在环
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
