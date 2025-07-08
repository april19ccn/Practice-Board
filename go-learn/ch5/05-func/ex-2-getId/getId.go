package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// go run .\getId.go >testLocal.html
	if len(os.Args) < 2 {
		doc, err := getHTML(true, "./template.html")
		if err != nil {
			fmt.Fprintf(os.Stderr, "getHTML: %v\n", err)
		}

		res := ElementByID(doc, "test")
		if res != nil {
			fmt.Println(res.Data)
		}
	} else {
		// go run .\getId.go http://gopl.io >testHttp.html
		for _, url := range os.Args[1:] {
			doc, err := getHTML(false, url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "getHTML: %v\n", err)
				continue
			}

			res := ElementByID(doc, "toc")
			if res != nil {
				fmt.Println(res.Data)
			}
		}
	}
}

// 获取HTML数据
func getHTML(isLocal bool, url string) (*html.Node, error) {
	var res io.Reader

	if isLocal {
		fileRes, err := os.ReadFile(url)
		if err != nil {
			if err == io.EOF {
				return nil, nil
			}
			return nil, err
		}
		res = bytes.NewReader(fileRes)
	} else {
		httpRes, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer httpRes.Body.Close()
		res = httpRes.Body
	}

	return html.Parse(res)
}

// 获取ID
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, endElement)
}

// !+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node := forEachNode(c, id, pre, post); node != nil {
			return node
		}
	}

	if post != nil {
		post(n, id)
	}

	return nil
}

//!-forEachNode

// !+startend
func startElement(n *html.Node, target string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == target {
				fmt.Println(attr.Val)
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node, target string) bool {
	return true
}

//!-startend
