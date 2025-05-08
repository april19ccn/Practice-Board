package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)
	out, err := os.Create("out.txt")

	if err != nil {
		return
	}

	for i := 0; i < 3; i++ {
		for _, url := range os.Args[1:] {
			go fetch(url, ch) // start a goroutine
		}
		for range os.Args[1:] {
			// fmt.Println(<-ch) // receive from channel ch
			out.WriteString(<-ch)
			out.WriteString("\n")
		}
		// fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
		out.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
		if i != 2 {
			out.WriteString("===============分 割 线===============\n")
		}
	}
	out.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body) // io.Discard 可以把这个变量看作一个垃圾桶，可以向里面写一些不需要的数据
	resp.Body.Close()                             // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url) // Sprintf 根据格式指定符进行格式化，并返回结果字符串。
}
