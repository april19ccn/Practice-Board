package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// soleTitle returns the text of the first non-empty title element
// in doc, and an error if there was not exactly one.
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil: // no panic
		case bailout{}: // "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()
	// Bail out of recursion if we find more than one nonempty title.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple titleelements
			}
			title = n.FirstChild.Data
			if title == "x" {
				fmt.Println("x---")
				panic("x")
			}
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func main() {
	htmlStr, err := os.ReadFile("test.html")
	if err != nil {
		fmt.Println(err)
	}
	doc, err := html.Parse(bytes.NewReader(htmlStr))
	if err != nil {
		fmt.Println(err)
	}
	title, err := soleTitle(doc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(title)
}
