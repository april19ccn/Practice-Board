package main

import (
	"context"
	newlinks "example/learn/ch8/09-exit/ex-01-cancel/newlinks"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
)

type EXData struct {
	depth int
	data  []string
}

// tokens is a counting semaphore used to enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(ctx context.Context, url string, depth int) EXData {
	fmt.Println(url)

	// 检查是否已取消
	select {
	case <-ctx.Done():
		return EXData{depth + 1, nil}
	default:
	}

	tokens <- struct{}{} // acquire a token
	list, err := newlinks.Extract(url, ctx)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return EXData{depth + 1, list}
}

func main() {
	// 添加defer捕获panic并打印所有goroutine栈
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic caught: %v\n", r)
			buf := make([]byte, 1<<20)
			stackSize := runtime.Stack(buf, true)
			fmt.Printf("=== ALL GOROUTINES ===\n%s\n", buf[:stackSize])
		}
	}()

	// 使用context来处理取消
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		cancel()                       // 取消所有操作
		fmt.Println("取消信号已发送")
	}()

	worklist := make(chan EXData, 100) // 增加缓冲区大小
	var wg sync.WaitGroup

	// 启动初始任务
	if len(os.Args) > 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case worklist <- EXData{0, os.Args[1:]}:
			case <-ctx.Done():
			}
		}()
	} else {
		fmt.Println("请提供起始URL作为参数")
		return
	}

	// 使用一个单独的goroutine来等待所有工作完成并关闭worklist
	go func() {
		wg.Wait()
		close(worklist)
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	// 使用for-select模式来处理worklist和取消信号
	for {
		select {
		case <-ctx.Done():
			// 上下文已取消，等待剩余工作完成
			fmt.Println("上下文已取消，等待剩余工作完成...")
			wg.Wait()
			return
		case t, ok := <-worklist:
			if !ok {
				// worklist已关闭，所有工作完成
				fmt.Println("所有工作完成")
				return
			}

			if t.depth < 2 {
				for _, link := range t.data {
					if !seen[link] {
						seen[link] = true
						wg.Add(1)
						go func(link string, depth int) {
							defer wg.Done()
							select {
							case worklist <- crawl(ctx, link, depth):
							case <-ctx.Done():
							}
						}(link, t.depth)
					}
				}
			} else {
				fmt.Println("大于2访问层级了!!")
			}
		}
	}
}
