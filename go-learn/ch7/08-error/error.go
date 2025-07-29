package main

import (
	"fmt"
	"syscall"
)

func main() {
	var err error = syscall.Errno(2)
	fmt.Println(err.Error()) // "no such file or directory"
	fmt.Println(err)
}

// 内置构建的 error 接口
// type error interface {
//     Error() string
// }

// errors包
// package errors

// func New(text string) error { return &errorString{text} }

// type errorString struct { text string }

// func (e *errorString) Error() string { return e.text }
