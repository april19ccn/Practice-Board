package main

import "testing"

var dataList = []struct {
	id      int
	graph   map[string][]string
	isCycle bool
}{
	{
		// 环  A-B—C
		id: 1,
		graph: map[string][]string{
			"A": {"B", "D"},
			"B": {"C"},
			"C": {"A"},
		},
		isCycle: true,
	},
	{
		// 环  A-B-C-D
		id: 2,
		graph: map[string][]string{
			"A": {"B"},
			"B": {"C"},
			"C": {"D"},
			"D": {"A"},
		},
		isCycle: true,
	},
	{
		// 无环
		id: 3,
		graph: map[string][]string{
			"A": {"B", "D"},
			"B": {"C"},
			"C": {"D"},
		},
		isCycle: false,
	},
	{
		// 无环
		id: 4,
		graph: map[string][]string{
			"A": {"B", "C", "D"},
			"B": {"C"},
			"C": {"D"},
		},
		isCycle: false,
	},
}

func TestFindCycle(t *testing.T) {
	for _, data := range dataList {
		got := checkCycle(data.graph)
		if got != data.isCycle {
			t.Errorf("id = %d, got = %v, want %v", data.id, got, data.isCycle)
		}
	}
}
