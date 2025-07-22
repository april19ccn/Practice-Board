// 练习 7.5： io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
// 并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。
// 实现这个LimitReader函数：
// func LimitReader(r io.Reader, n int64) io.Reader

package main

import "io"

// 定义结构体，包含底层 Reader 和剩余可读字节数
type LimitedReader struct {
	R io.Reader
	N int64
}

// Read 方法实现 io.Reader 接口
func (l *LimitedReader) Read(p []byte) (n int, err error) {
	// 如果剩余字节数 <= 0，直接返回 EOF
	if l.N <= 0 {
		return 0, io.EOF
	}
	// 调整读取长度不超过剩余字节数
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	// 从底层 Reader 读取数据
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

// LimitReader 返回一个最多读取 n 字节的 Reader
func LimitReader(r io.Reader, n int64) io.Reader { return &LimitedReader{r, n} }

func main() {
	io.LimitReader(nil, 0)

	// io.Reader(nil)
}
