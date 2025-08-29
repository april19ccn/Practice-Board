// Package links provides a link-extraction function.
package newlinks

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

// func main() {
// 	for _, url := range os.Args[1:] {
// 		links, err := Extract(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
// 			continue
// 		}
// 		fmt.Fprintf(os.Stdout, "%s: %d links\n", url, len(links))
// 		for _, link := range links {
// 			fmt.Println(link)
// 		}
// 	}
// }

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
// The request can be cancelled by closing the cancel channel.
func Extract(url string, ctx context.Context) ([]string, error) {
	tr := &http.Transport{
		IdleConnTimeout:   1 * time.Second,
		DisableKeepAlives: true, // 禁用持久连接
	}
	client := &http.Client{
		Transport: tr,
	}

	// 创建一个新的请求，而不是使用http.Get
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	// // 设置Cancel字段以支持取消请求
	// req.Cancel = cancel

	// 使用http.DefaultClient.Do发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val) // 现在links中存储的不是href属性的原始值，而是通过resp.Request.URL解析后的值。
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
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
