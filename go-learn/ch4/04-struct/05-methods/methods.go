package main

import "fmt"

type Counter int

func (c *Counter) Increment() { *c++ }
func (c Counter) get() int    { return int(c) } // 注意：小写开头的私有方法

type Container struct {
	Counter
	Name string
}

func main() {
	c := Container{Name: "Demo"}
	c.Increment() // ✅ 仍然有效
	c.Increment()
	fmt.Println(c.get())
}
