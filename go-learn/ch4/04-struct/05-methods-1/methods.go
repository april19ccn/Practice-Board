package main

import "fmt"

type Base struct {
	age int
}

func (b *Base) sayHello() {
	fmt.Printf("Hello! I'm %v years old!", b.age)
}

type Derived struct {
	Base
}

func main() {
	a := Derived{Base{30}}
	fmt.Println(a.age)
	a.sayHello()
}
