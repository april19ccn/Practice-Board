package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// ******* io.Writer 接口
type UpperWriter struct {
	io.Writer
}

// 这里不要与接口内嵌弄混了
// 接口内嵌：
// [接口类型 - Go语言圣经](https://golang-china.github.io/gopl-zh/ch7/ch7-02.html#WCREFX-12044696)
// 上面用到的语法和结构内嵌相似，我们可以用这种方式以一个简写命名一个接口，而不用声明它所有的方法。这种方式称为接口内嵌。

// type UpperWriter struct {
// 	Writer io.Writer
// }
// 这里是内嵌了一个 io.Writer 接口类型的字段
// &UpperWriter{os.Stdout}
// => UpperWriter{Writer: os.Stdout}

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

type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
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

	// var (
	// 	a io.ReadCloser = (*os.File)(f) // 隐式转换，*os.File满足io.ReadCloser接口
	// 	b io.Reader     = a             // 隐式转换，io.ReadCloser满足io.Reader接口
	// 	c io.Closer     = a             // 隐式转换，io.ReadCloser满足io.Closer接口
	// 	d io.Reader     = c.(io.Reader) // 显式转换，io.Closer不满足io.Reader接口
	// )

	var tb testing.TB = new(TB)
	tb.Fatal("Hello, playground")
	// 为什么 tb 不应该是 *testing.TB？
	// 你疑惑的点可能是：“既然右边是指针，左边不也应该是指针吗？”

	// 如果写成 *testing.TB，它的含义是 “指向接口的指针” 。在 Go 语言中，几乎永远不需要（也不应该）使用指向接口的指针。

	// 接口本质上已经像是一个指针：接口类型的底层实现包含两个指针，一个指向类型信息，一个指向具体的数据。
	// 多此一举：如果你定义 var tb *testing.TB，你就得去取接口的地址。这不仅冗余，而且会导致很多令人困惑的行为。
}
