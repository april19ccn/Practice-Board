// 验证接口：
// T 类型的方法集：只有 ValueMethod（不包含 PtrMethod）
// *T 类型的方法集：包含 ValueMethod 和 PtrMethod
package main

import "fmt"

type T struct{}

// 值接收者方法
func (t T) ValueMethod() {}

// 指针接收者方法
func (t *T) PtrMethod() {}

// 定义两个接口
type Interface1 interface {
	ValueMethod()
}

type Interface2 interface {
	PtrMethod()
}

// 组合一个接口
type Interface3 interface {
	ValueMethod()
	PtrMethod()
}

func main() {
	var i1 Interface1
	var i2 Interface2
	var i3 Interface3

	i1 = T{}
	i1 = &T{}
	// i2 = T{} // 不能在赋值中使用 T{}（结构类型 T 的值）作为 Interface2 值：T 不实现 Interface2（方法 PtrMethod 具有指针接收器）
	i2 = &T{}
	// i3 = T{} // cannot use T{} (value of struct type T) as Interface3 value in assignment: T does not implement Interface3 (method PtrMethod has pointer receiver)
	i3 = &T{}

	fmt.Println(i1, i2, i3) // 消除 i1、i2、i3 提示未使用的错误
}
