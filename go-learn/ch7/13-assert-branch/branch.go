package main

import (
	"fmt"
	"io"
)

func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x has type interface{} here.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return fmt.Sprintf("string: %s", x) // (not shown)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

type TestInt int

func (t TestInt) String() string {
	return fmt.Sprintf("int: %d", t)
}

func main() {
	fmt.Println(sqlQuote(42))             // 42
	fmt.Println(sqlQuote("hello"))        // string: hello
	fmt.Println(sqlQuote(true))           // TRUE
	fmt.Println(sqlQuote(nil))            // NULL
	fmt.Println(sqlQuote(io.Writer(nil))) // NULL

	// fmt.Println(sqlQuote(TestInt(42))) // panic: unexpected type main.TestInt: int: 42
	// fmt.Println(sqlQuote(*os.Stdin))   // unexpected type os.File: {0xc0000a0008}
}
