package main

import (
	inset "example/learn/ch6/05-bit/ex-inset"
)

func main() {
	var got = inset.IntSet{}
	got.Add(1)
	got.Add(55)
	got.Add(64)
	got.Add(188)
	// fmt.Println(got.words) // 小写不可见
}
