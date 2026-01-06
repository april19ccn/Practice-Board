package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// ******* io.Writer 接口
type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))

	// 为什么不能省略 Writer ？
	// 执行链：Write -> 调用 p.Write -> 再次进入 Write -> 再次调用 p.Write -> ...
	// 结果：函数无限调用自己，导致栈溢出，程序崩溃。
	// return p.Write(bytes.ToUpper(data))

	// 为什么 p.Writer 就是 UpperWriter 里的 io.Writer ？
	// [结构体 - Go语言圣经](https://golang-china.github.io/gopl-zh/ch4/ch4-04.html#WCREFX-12015571)
	// 其中匿名成员Circle和Point都有自己的名字——就是命名的类型名字——但是这些名字在点操作符中是可选的。我们在访问子成员的时候可以忽略任何匿名成员部分。
}

// ******* fmt.Stringer 接口
// type fmt.Stringer interface {
//     String() string
// }

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

func main() {
	// fmt.Printf("hello")
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")

	fmt.Fprintln(os.Stdout, UpperString("hello, world"))

	// go 1.24.2
	// handleMethods 源码里 需要判断 UpperString 是否实现了 fmt.Stringer 接口
	// // Is it an error or Stringer?
	// // The duplication in the bodies is necessary:
	// // setting handled and deferring catchPanic
	// // must happen before calling the method.
	// switch v := p.arg.(type) {
	// case error:
	// 	handled = true
	// 	defer p.catchPanic(p.arg, verb, "Error")
	// 	p.fmtString(v.Error(), verb)
	// 	return

	// case Stringer: // 判断 UpperString 符合 fmt.Stringer 接口
	// 	handled = true
	// 	defer p.catchPanic(p.arg, verb, "String")
	// 	p.fmtString(v.String(), verb)
	// 	return
	// }
}
