// 练习 8.11： 紧接着8.4.4中的mirroredQuery流程，实现一个并发请求url的fetch的变种。当第一个请求返回时，直接取消其它的请求。
package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

func mirroredQuery(ctx context.Context) chan string {
	var wg sync.WaitGroup
	wg.Add(3)

	responses := make(chan string, 3)
	go func() {
		defer wg.Done()
		resp, err := request("https://www.baidu.com/", ctx)
		if err != nil {
			fmt.Println(err)
		} else {
			responses <- resp
		}
	}()
	go func() {
		defer wg.Done()
		resp, err := request("https://www.npmjs.com/", ctx)
		if err != nil {
			fmt.Println(err)
		} else {
			responses <- resp
		}
	}()
	go func() {
		defer wg.Done()
		resp, err := request("https://www.google.com.hk/", ctx)
		if err != nil {
			fmt.Println(err)
		} else {
			responses <- resp
		}
	}()
	// 等待所有goroutine完成，然后关闭通道
	// go func() {
	// 	wg.Wait()
	// 	close(responses)
	// }()
	return responses
}

func request(hostname string, ctx context.Context) (response string, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", hostname, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}
	return hostname, err
}

// func request(hostname string, ctx context.Context) (response string, err error) {
// 	return hostname, err
// }

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range mirroredQuery(ctx) {
		fmt.Println(v)
		cancel()
	}
}
