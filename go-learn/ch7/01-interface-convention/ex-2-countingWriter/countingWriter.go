// 练习 7.2： 写一个带有如下函数签名的函数CountingWriter，
// 传入一个io.Writer接口类型，
// 返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。
// func CountingWriter(w io.Writer) (io.Writer, *int64)
package main

import (
	"fmt"
	"io"
	"os"
)

type byteCounter struct {
	writer  io.Writer
	counter *int64
}

func (c *byteCounter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p) // c自动解引用
	*(*c).counter += int64(n)   // 手动解引用
	return n, err
}

func (c *byteCounter) Counter() *int64 {
	return c.counter
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	b := &byteCounter{w, new(int64)}
	return b, b.Counter()
}

func main() {
	b, counts := CountingWriter(os.Stdout)
	b.Write([]byte("hello, world\n"))
	fmt.Println(*counts)
}
