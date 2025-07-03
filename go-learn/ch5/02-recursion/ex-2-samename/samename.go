package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func statistics(stats map[string]int, n *html.Node) {
	// 如果当前节点为空，递归结束, 如果下面是 for 循环，可以不需要这个条件判断
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		stats[n.Data]++
	}
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	statistics(stats, c)
	// }

	statistics(stats, n.FirstChild)
	statistics(stats, n.NextSibling)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	counts := make(map[string]int)
	statistics(counts, doc)

	for k, v := range counts {
		fmt.Printf("%s: %d\n", k, v)
	}
}
