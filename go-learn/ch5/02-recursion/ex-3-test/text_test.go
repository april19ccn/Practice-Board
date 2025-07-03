package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestText(t *testing.T) {
	htmlStr, err := os.ReadFile("test.html")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := html.Parse(bytes.NewReader(htmlStr))
	if err != nil {
		t.Fatal(err)
	}
	texts := text(nil, doc)
	fmt.Println(texts)
}
