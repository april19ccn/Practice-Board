package main

import "testing"

func TestTopoSort(t *testing.T) {
	got := topoSort(Prereqs)

	for k, v := range Prereqs {
		target := got[k]
		for _, item := range v {
			if got[item] >= target { // 利用拓扑排序后，前置条件的课程必然小于当前课程
				t.Errorf("got %v: %v, target %v: %v", item, got[item], k, target)
			}
		}
	}
}
