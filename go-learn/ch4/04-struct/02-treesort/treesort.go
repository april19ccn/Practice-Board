package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// 将数值就地排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues 将 t 的元素按顺序追加到值中。
// 并返回结果切片。
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// add 将值 value 添加到树 t 中。
func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value // 语法糖，自动解引用 (*t).value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	test := []int{5, 2, 6, 3, 1, 4, 10, 7}
	Sort(test)
	fmt.Println(test)
}
