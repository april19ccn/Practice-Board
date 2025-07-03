// 练习 5.3： 编写函数输出所有text结点的内容。
// 注意不要访问<script>和<style>元素，因为这些元素对浏览者是不可见的。
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func text(stack []string, n *html.Node) []string {
	if n == nil || n.Data == "script" || n.Data == "style" {
		return stack
	}

	if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
		stack = append(stack, n.Data)
	}

	stack = text(stack, n.FirstChild)
	stack = text(stack, n.NextSibling)

	return stack
}

func main() {
	doc, err := html.Parse(os.Stdin)
	fmt.Println(doc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "text: %v\n", err)
		os.Exit(1)
	}
	for _, link := range text(nil, doc) {
		fmt.Println(link)
	}
}

// ./fetch https://golang.org | .\text.exe
