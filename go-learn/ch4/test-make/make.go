package main

func testMake() []int {
	arr := make([]int, 6) // 0.2658 ns/op

	s := []int{0, 1, 2, 3, 4, 5}
	for index, v := range s {
		arr[v] = s[len(s)-1-index]
	} // 2.949 ns/op

	return arr
}
