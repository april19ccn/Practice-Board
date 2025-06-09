package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// type Weekday1 string

// const (
// 	Sunday1 Weekday1 = "星期" + string(iota+1) // Go 要求常量必须在编译时完全确定值，复杂的字符串拼接可能涉及运行时行为。
// 	Monday1
// 	Tuesday1
// 	Wednesday1
// 	Thursday1
// 	Friday1
// 	Saturday1
// )

func main() {
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
