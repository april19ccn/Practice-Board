package main

import "testing"

func TestTopoSort(t *testing.T) {
	if checkCycle(Prereqs) {
		t.Log("发现环，但测试通过") // 记录信息，继续执行
		return             // 提前返回，结束测试
	}
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
