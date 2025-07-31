// 仅用于总结 类型断言 这一章，并非严谨可执行代码
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Reader interface{ Read() }
type Writer interface{ Write() }
type ReadWriter interface {
	Reader
	Writer
}

// 实现类型
type File struct{}

func (f File) Read()  {}
func (f File) Write() {}

func main() {
	// 具体类型 断言 具体类型
	// 证明 GO的类型断言 只能用于接口
	// buf1 := new(bytes.Buffer)
	// buf2 := new(bytes.Buffer)
	// buf1.(buf2) // invalid operation: buf1 (variable of type *bytes.Buffer) is not an interface
	// => 验证类型断言只能用于接口

	// 具体类型 断言 接口类型 (并不可以使用)
	// 证明 w.(T) 的 w 必须是接口类型
	// var file File
	// filew, _ := file.(Writer) // invalid operation: file (variable of struct type File) is not an interface

	// 综上 类型断言只有两种：

	// 1. 接口类型 断言 具体类型
	// 类型断言 w.(*os.File) 尝试验证接口值 w 底层是否持有 *os.File 类型，如果验证成功，则返回该类型对应的具体值。
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)      // success: f == os.Stdout
	c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer

	fmt.Println(f, c)

	// 2. 接口类型 断言 接口类型 （提升到限制更多的接口）
	var r Reader = File{} // r 的类型为 Reader, 动态类型 File

	// 断言为 Writer 接口
	fw, ok := r.(Writer)
	// ✔ 断言成功：因为 File 实现了 Writer
	// w 的类型变为 Writer, 但底层仍是 File 值

	if ok { // 利用ok，避免panic
		fw.Write() // 现在可调用 Writer 的方法
	}

	fw, _ = r.(ReadWriter) // fw 的类型 已经在第一次声明了，所以这里并不会修改fw的类型， 且  Go 语言的类型系统特性：当您将值赋给已声明类型的变量时，该值会被转换为目标变量的类型。
	// fw.Read() // fw.Read undefined (type Writer has no field or method Read)
	fw.Write() // 现在可调用 Writer 的方法

	// 更多限制的接口类型 是没必要断言 更少限制的接口类型
	var rw ReadWriter = File{}
	// filew, _ := rw.(Writer)
	var filew Writer = rw // 等价直接赋值
	fmt.Println(filew)

}
