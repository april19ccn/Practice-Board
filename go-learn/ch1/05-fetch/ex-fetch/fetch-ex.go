// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// 练习1.8
		if !strings.HasPrefix(url, "http://") { // HasPrefix 报告字符串 s 是否以前缀开头
			url += "http://"
		}

		// http.Get函数是创建HTTP请求的函数
		// resp的Body字段包括一个可读的服务器响应流
		resp, err := http.Get(url)
		fmt.Println(resp.Status) // 练习1.9 resp.Status变量得到该状态码
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body) // 练习1.7
		resp.Body.Close()                      // resp.Body.Close关闭resp的Body流，防止资源泄露
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
