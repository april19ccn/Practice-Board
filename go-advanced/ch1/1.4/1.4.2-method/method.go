package main

import "fmt"

// 文件对象
type File struct {
	fd int
}

// 打开文件
func OpenFile(name string) (f *File, err error) {
	// ...
	return nil, nil
}

// 关闭文件
func (f *File) Close() error {
	// ...
	return nil
}

// 读文件数据
func (f *File) Read(offset int64, data []byte) int {
	// ...
	return 0
}

type Point struct{}

func (p Point) B() { fmt.Println("Point B") }
func (p Point) A() { p.B() } // Point 的 A 调用 B

type ColoredPoint struct{ Point }

func (c ColoredPoint) B() { fmt.Println("ColoredPoint B") } // ColoredPoint 实现了自己的 B

func main() {
	// 不依赖具体的文件对象
	// func CloseFile(f *File) error
	var CloseFile = (*File).Close

	// 不依赖具体的文件对象
	// func ReadFile(f *File, offset int64, data []byte) int
	var ReadFile = (*File).Read

	// 文件处理
	f, _ := OpenFile("foo.dat")
	data := []byte{}
	ReadFile(f, 0, data)
	CloseFile(f)

	var cp ColoredPoint
	cp.A()
	cp.B()
	cp.Point.B()
}
