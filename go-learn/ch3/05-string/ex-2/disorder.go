package main

import (
	"fmt"
	"sort"
	"strings"

	S "github.com/IBM/fp-go/string"
)

// 字符串->数组->排序->字符串
func disorderArr(left string, right string) bool {
	if S.Eq(left, right) || len(left) != len(right) {
		return false
	}

	leftArr := strings.Split(left, "")
	rightArr := strings.Split(right, "")

	sort.Strings(leftArr)
	sort.Strings(rightArr)

	return strings.Join(leftArr, "") == strings.Join(rightArr, "")
}

// 哈希表
func disorderMap(left string, right string) bool {
	if S.Eq(left, right) || len(left) != len(right) {
		return false
	}

	strMap := make(map[rune]int)

	for _, v := range left {
		strMap[v]++
	}

	for _, v := range right {
		strMap[v]--
		if strMap[v] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(disorderArr("abc123中文", "文123abc中"))
}
