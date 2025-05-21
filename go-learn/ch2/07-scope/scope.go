package main

import (
	"fmt"
	"log"
	"os"
)

func f() {}

var g = "g"

var cwd string

// test5
func init() {
	cwd, err := os.Getwd() // NOTE: wrong!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	fmt.Println(cwd)
	// log.Printf("Working directory = %s", cwd)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}

func test1() {
	f := "f"
	fmt.Println(f) // "f"; local var f shadows package-level func f
	fmt.Println(g) // "g"; package-level var
	// fmt.Println(h) // compile error: undefined: h
}

func test2() {
	// x 不同作用域 （只是示范，实际不推荐这样做）
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}
	fmt.Print("\n")
}

func test3() {
	x := "hello"          // 声明 x 在函数体词法域
	for _, x := range x { // 声明 x' 在for隐式的初始化词法域
		x := x + 'A' - 'a'  // 声明 x'' 在for循环体词法域
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
	fmt.Print("\n")
}

func test4() {
	var fx = func() int { return 2 }
	var gx = func(x int) int { return x + 1 }

	if x := fx(); x == 0 {
		z := 3
		fmt.Println(x)
		fmt.Println(z)
	} else if y := gx(x); x == y {
		fmt.Println(x, y)
	} else {
		// fmt.Println(z)  // undefined: z 条件部分为一个隐式词法域 各个分支都可以调用，但每个分支的词法域是隔离的
		fmt.Println(x, y)
	}

	// fmt.Println(x, y) // compile error: x and y are not visible here

	fmt.Print("\n")
}
