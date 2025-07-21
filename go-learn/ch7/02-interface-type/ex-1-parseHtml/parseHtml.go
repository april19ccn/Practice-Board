// 练习 7.4： strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。
// 实现一个简单版本的NewReader，用它来构造一个接收字符串输入的HTML解析器（§5.2）
package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type ParseHtml struct {
	html string
	i    int64 // current reading index
}

func (p *ParseHtml) Read(b []byte) (n int, err error) {
	if p.i >= int64(len(p.html)) {
		return 0, io.EOF
	}
	// 在Go语言中，copy函数可以处理从字符串（string）到字节切片（[]byte）的复制。
	// 这是因为字符串底层就是字节切片（只读），所以copy函数能够安全地将字符串的字节内容复制到目标字节切片中。
	n = copy(b, p.html[p.i:])
	p.i += int64(n)
	return
}

func NewReader(s string) *ParseHtml { return &ParseHtml{s, 0} }

func main() {
	strings.NewReader("hello")
	htmlStr, err := os.ReadFile("test.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	doc, err := html.Parse(NewReader(string(htmlStr)))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v", doc)
}
