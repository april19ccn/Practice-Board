//go:build amd64 || arm64
// +build amd64 arm64

package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
	"unsafe"
)

func initSlice() {
	var (
		a []int               // nil切片，和nil相等，一般用来表示一个不存在的切片
		b = []int{}           // 空切片，和nil不相等，一般用来表示一个空的集合
		c = []int{1, 2, 3}    // 有3个元素的切片，len和cap都为3
		d = c[:2]             // 有2个元素的切片，len为2, cap为3
		e = c[0:2:cap(c)]     // 有2个元素的切片，len为2, cap为3
		f = c[:0]             // 有0个元素的切片，len为0, cap为3
		g = make([]int, 3)    // 有3个元素的切片，len和cap都为3
		h = make([]int, 2, 3) // 有2个元素的切片，len为2, cap为3
		i = make([]int, 0, 3) // 有0个元素的切片，len为0, cap为3
	)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println("-----")
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
}

// 探究数组复制和切片复制的差异
// 切片复制的结构体
// 数组是整个值重新生成（验证如下）
func arrayCopy() {
	test := [5]int{1, 2, 3, 4, 5}
	x := test
	fmt.Printf("test value: %p\n", &test)
	fmt.Printf("x value: %p\n", &x)
	x[4] = 6
	fmt.Println(test)
}

// 1 中间添加切片元素
func AddMid() {
	a := []int{1, 2, 3, 4, 5}
	x := []int{0, 0, 0}
	a = append(a, x...)       // [1 2 3 4 5 0 0 0] // append 扩展切片只是副作用
	copy(a[1+len(x):], a[1:]) // [5 0 0 0] <- [2 3 4 5] => [1 2 3 4 2 3 4 5]
	fmt.Println(a)            // [1 2 3 4 2 3 4 5]
	copy(a[1:], x)            // [2 3 4 2 3 4 5] <- [0 0 0] => [0 0 0 2 3 4 5]
	fmt.Println(a)            // [1 0 0 0 2 3 4 5]
}

// 2 中间删除切片元素
func DelMid() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{9, 8, 7, 6, 7, 4, 3}

	i := 1
	N := 3

	// a = append(a[:i], a[i+1:]...) // 删除中间1个元素 [1] + [3, 4, 5, 6]
	a = append(a[:i], a[i+N:]...) // 删除中间N个元素 [1] + [5, 6]
	fmt.Println(a)

	b = b[:i+copy(b[i:], b[i+1:])] // 删除中间1个元素
	// copy([8, 7, 6, 7, 4, 3], [7, 6, 7, 4, 3]) => [7, 6, 7, 4, 3, 3]
	// [:i+5] => [9, 7, 6, 7, 4, 3]
	// b = b[:i+copy(b[i:], b[i+N:])] // 删除中间N个元素
	fmt.Println(b)
}

// 3 切片内存利用
func TrimSpace(s []byte) []byte {
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}

func Filter(s []byte, fn func(x byte) bool) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}

// 4
func FindPhoneNumber(filename string) []byte {
	b, _ := os.ReadFile(filename)
	return regexp.MustCompile("[0-9]+").Find(b)
}

// 5 切片类型强制转换
var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}

func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以int方式给float64排序（不支持负数）
	sort.Ints(b)
}

func SortFloat64FastV2(a []float64) {
	// 通过reflect.SliceHeader更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr

	// 以int方式给float64排序（不支持负数）
	sort.Ints(c)
}

func SortFloat64FastV3(a []float64) {
	c := unsafe.Slice(
		(*int)(unsafe.Pointer(unsafe.SliceData(a))),
		len(a),
	)

	// 以int方式给float64排序（不支持负数）
	sort.Ints(c)
}

func main() {
	// initSlice()
	// arrayCopy()
	// AddMid()
	// DelMid()

	// test := []byte{'h', ' ', ' ', ' ', ' ', 'e', ' ', 'l', 'l', ' ', 'o', ',', ' ', 'w', ' ', 'o', ' ', 'r', 'l', 'd', '.', ' '}
	// trimString := TrimSpace([]byte(test))
	// fmt.Printf("%s\n", trimString)
	// fmt.Printf("%s\n", test)

}
