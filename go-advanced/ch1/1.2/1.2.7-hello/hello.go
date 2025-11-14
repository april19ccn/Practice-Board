package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Please visit http://127.0.0.1:12345/")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		s := fmt.Sprintf("你好, 世界! -- Time: %s", time.Now().String())
		fmt.Fprintf(w, "%v\n", s)
		log.Printf("路径: %s, %v\n", req.URL.Path, s)
	})
	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 为什么在访问网站后，控制台会打印两条
// 2025/11/14 15:56:55 你好, 世界! -- Time: 2025-11-14 15:56:55.233204 +0800 CST m=+5.194789709
// 2025/11/14 15:56:55 你好, 世界! -- Time: 2025-11-14 15:56:55.273204 +0800 CST m=+5.234789293

// 这个现象是因为浏览器在访问网页时通常会发送两个请求：
// - 对主页的请求 (GET /)
// - 对 favicon.ico 的请求 (GET /favicon.ico)
