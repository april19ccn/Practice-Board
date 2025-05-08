// ps: deepseek 给的答案
package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	// 第一次请求
	filename1, secs1, nbytes1, err := fetchAndSave(url, 1)
	if err != nil {
		ch <- fmt.Sprintf("第一次请求 %s 失败: %v", url, err)
		return
	}

	// 第二次请求
	filename2, secs2, nbytes2, err := fetchAndSave(url, 2)
	if err != nil {
		ch <- fmt.Sprintf("第二次请求 %s 失败: %v", url, err)
		return
	}

	// 比较内容
	same, err := compareFiles(filename1, filename2)
	if err != nil {
		ch <- fmt.Sprintf("比较文件失败: %v", err)
		return
	}

	ch <- fmt.Sprintf(
		"%.2fs (第一次) → %.2fs (第二次), 差异: %.2fs, 字节数: %d → %d, 内容一致: %t  %s",
		secs1, secs2, secs2-secs1, nbytes1, nbytes2, same, url,
	)
}

func fetchAndSave(url string, num int) (string, float64, int, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		return "", 0, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, 0, err
	}
	secs := time.Since(start).Seconds()
	nbytes := len(body)

	filename := generateFilename(url, num)
	if err := os.WriteFile(filename, body, 0644); err != nil {
		return "", 0, 0, err
	}

	return filename, secs, nbytes, nil
}

func generateFilename(url string, num int) string {
	hash := md5.Sum([]byte(url))
	return fmt.Sprintf("%x_%d.html", hash, num)
}

func compareFiles(file1, file2 string) (bool, error) {
	c1, err := os.ReadFile(file1)
	if err != nil {
		return false, err
	}
	c2, err := os.ReadFile(file2)
	return bytes.Equal(c1, c2), err
}
