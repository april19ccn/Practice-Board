package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	fmt.Println(doc)
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.ElementNode && (n.Data == "img") {
		images++
	}
	if n.Type == html.TextNode {
		// input := bufio.NewScanner(bytes.NewReader([]byte(n.Data)))
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)

		for input.Scan() {
			word := input.Text()
			fmt.Println(word)
			words++
		}
	}
	if n.FirstChild != nil {
		w, i := countWordsAndImages(n.FirstChild)
		words += w
		images += i
	}
	if n.NextSibling != nil {
		w, i := countWordsAndImages(n.NextSibling)
		words += w
		images += i
	}
	return
}

func main() {
	words, images, err := CountWordsAndImages("http://127.0.0.1:8080/test.html")
	// words, images, err := CountWordsAndImages("http://golang.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("words: %d, images: %d\n", words, images)
}
