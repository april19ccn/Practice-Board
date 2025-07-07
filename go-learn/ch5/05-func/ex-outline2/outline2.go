// 练习 5.7

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// go run .\outline2.go >testLocal.html
	if len(os.Args) < 2 {
		outline(os.Stdout, true, "./template.html")
	}

	// go run .\outline2.go http://gopl.io >testHttp.html
	for _, url := range os.Args[1:] {
		outline(os.Stdout, false, url)
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

func outline(w io.Writer, isLocal bool, url string) error {
	doc, err := getHTML(isLocal, url)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(w, doc, startElement, endElement)
	//!-call

	return nil
}

// !+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(w io.Writer, n *html.Node, pre, post func(w io.Writer, n *html.Node)) {
	if pre != nil {
		pre(w, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(w, c, pre, post)
	}

	if post != nil {
		post(w, n)
	}
}

//!-forEachNode

// !+startend
var depth int
var selfTag = map[string]bool{"img": true, "br": true, "link": true, "meta": true}

func startElement(w io.Writer, n *html.Node) {
	// fmt.Println(n.Data)
	if n.Type == html.ElementNode {
		// 构造标签和属性
		res := n.Data + " "
		for _, attr := range n.Attr {
			res += attr.Key + "=\"" + attr.Val + "\" "
		}
		if selfTag[n.Data] {
			fmt.Fprintf(w, "%*s<%s/>\n", depth*2, "", res)
		} else {
			fmt.Fprintf(w, "%*s<%s>\n", depth*2, "", res)
		}
		depth++
	}
	if n.Type == html.CommentNode {
		fmt.Fprintf(w, "%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
	if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
		fmt.Fprintf(w, "%*s%s\n", depth*2+2, "", strings.TrimSpace(n.Data))
	}
}

func endElement(w io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if !selfTag[n.Data] {
			fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

//!-startend
