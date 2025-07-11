package main

import (
	"fmt"
	"reflect"
	"testing"
)

var dataList = []struct {
	id      int
	graph   map[string][]string
	isCycle bool
	num     int
	want    [][]string
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
		num:     1,
		want:    [][]string{{"A", "B", "C", "A"}, {"B", "C", "A", "B"}, {"C", "A", "B", "C"}},
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
		num:     1,
		want:    [][]string{{"A", "B", "C", "D", "A"}, {"B", "C", "D", "A", "B"}, {"C", "D", "A", "B", "C"}, {"D", "A", "B", "C", "D"}},
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
		num:     0,
		want:    [][]string{},
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
		num:     0,
		want:    [][]string{},
	},
}

func TestFindCycle(t *testing.T) {
	for _, data := range dataList {
		got := findCycle(data.graph)

		fmt.Println("got:", got, "want:", data.want)

		for i, v := range data.want {
			if !reflect.DeepEqual(got[i], v) {
				t.Errorf("id= %d, got[i] = %v, want %v", data.id, got[i], v)
			}
		}
	}
}
