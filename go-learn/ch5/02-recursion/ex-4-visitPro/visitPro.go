package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "script" || n.Data == "img" || n.Data == "link") {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "visitPro: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
