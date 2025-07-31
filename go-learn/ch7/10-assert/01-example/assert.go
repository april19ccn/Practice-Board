package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// var w io.Writer
	// w = os.Stdout
	// f := w.(*os.File)      // success: f == os.Stdout
	// c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer

	// fmt.Println(f, c)

	var w io.Writer
	// 1. 对 nil 接口值的断言总是失败
	f, ok := w.(*os.File)      // ok = false, 失败。断言具体类型失败。
	e, ok := w.(io.ReadWriter) // ok = false, 失败。断言更多限制接口失败。
	a, ok := w.(io.Writer)     // ok = false, 失败！断言自身类型也失败。
	b, ok := w.(interface{})   // ok = false, 失败！断言空接口也失败！

	fmt.Println(f, e, a, b, ok)

	// 2. 对非 nil 接口值断言到更少限制接口 (几乎总是成功，类似赋值)
	buf := new(bytes.Buffer) // *bytes.Buffer 实现了 io.Writer
	w = buf                  // 赋值，w 现在持有 (*bytes.Buffer, buf)

	// 断言到更少限制接口：io.Writer 的方法集 (Write) 是 io.ReadWriter 方法集 (Read, Write) 的子集吗？
	// 不！io.ReadWriter 要求更多方法 (Read)。所以这是断言到 *更多* 限制接口，可能会失败。
	rw, ok := w.(io.ReadWriter) // ok = true? 取决于 *bytes.Buffer 是否实现了 Read（它实现了，所以ok=true）

	fmt.Println(rw, ok)

	// 正确例子：断言到更少限制接口 (几乎总是成功)
	var rw2 io.ReadWriter = buf // rwc 要求 Read, Write, Close
	// 断言到 io.Writer (要求 Write)。io.Writer 的方法集是 io.ReadWriteCloser 方法集的子集。
	w2, ok := rw2.(io.Writer) // ok 必然为 true (只要 rwc 非nil)。行为等价于 `var w2 io.Writer = rwc`
	// 通常直接写赋值更好: w2 := rwc

	fmt.Println(w2, ok)

	// 3. "除了对于nil接口值的情况" - 结合点
	var nilRwc io.ReadWriteCloser // nil 接口值 (io.ReadWriteCloser=nil)
	// 断言到限制更少的接口 io.Writer (方法更少)
	nilW, ok := nilRwc.(io.Writer) // ok = false! 失败！尽管 io.Writer 限制更少，但原始接口是 nil。
	// 对比赋值：
	var w3 io.Writer = nilRwc // w3 现在是一个 nil 接口值 (io.Writer=nil)

	fmt.Println(nilW, w3)
}
