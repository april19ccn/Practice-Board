// 练习5.17： 编写多参数版本的ElementsByTagName，
// 函数接收一个HTML结点树以及任意数量的标签名，返回与这些标签名匹配的所有元素。
// 下面给出了2个例子：
package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// images := ElementsByTagName(doc, "img")
// headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var checkMap = make(map[string]bool, len(name))
	for _, v := range name {
		checkMap[v] = true
	}

	var nodes []*html.Node
	var check func(*html.Node)
	check = func(n *html.Node) {
		if n.Type == html.ElementNode && checkMap[n.Data] {
			fmt.Println(n.Data)
			nodes = append(nodes, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			check(c)
		}
	}
	check(doc)

	return nodes
}

func main() {
	htmlStr, err := os.ReadFile("test.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	doc, err := html.Parse(bytes.NewReader(htmlStr))
	if err != nil {
		fmt.Println(err)
		return
	}

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	fmt.Println(images, headings)
}
