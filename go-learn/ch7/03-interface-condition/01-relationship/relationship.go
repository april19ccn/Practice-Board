package main

import (
	"fmt"
	"reflect"
	"strings"
)

// 12.8章包含了一个打印出任意值的所有方法的程序(暂时搬过来用一下，具体原理等12章在研究)
// Print prints the method set of the value x.
func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}

type IntSet struct { /* ... */
}

func (*IntSet) String() string {
	return "IntSet"
}

func (*IntSet) String1() string {
	return "IntSet"
}

func (IntSet) String2() string {
	return "IntSet"
}

func main() {
	// IntSet类型的String方法的接收者是一个指针类型，所以我们不能在一个不能寻址的IntSet值上调用这个方法
	// var _ = IntSet{}.String() // compile error: String requires *IntSet receiver

	// 可以在一个IntSet变量上调用这个方法
	var s IntSet
	var _ = s.String() // OK: s is a variable and &s has a String method

	// 由于只有*IntSet类型有String方法，所以也只有*IntSet类型实现了fmt.Stringer接口
	var _ fmt.Stringer = &s // OK
	// var _ fmt.Stringer = s  // compile error: IntSet lacks String method

	Print(&s)
	// type *main.IntSet
	// func (*main.IntSet) String() string
	// func (*main.IntSet) String1() string
	// func (*main.IntSet) String2() string

	Print(s)
	// type main.IntSet
	// func (main.IntSet) String2() string
}
