package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "获取文件信息失败: %v\n", err)
				continue
			}
			fileSizes <- info.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.DirEntry {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "读取目录失败: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func scanDir(roots []string, lock *bool) {
	// Traverse the file tree.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
		mu.Lock()
		*lock = false
		mu.Unlock()
	}()

	//!+
	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

var verbose = flag.Bool("v", false, "show verbose progress messages")
var mu sync.Mutex

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	mu.Lock()
	isLock := true
	mu.Unlock()

	go scanDir(roots, &isLock)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		mu.Lock()
		scanning := isLock
		mu.Unlock()
		if scanning { // 确保读取到的是最新状态
			fmt.Println("当前正在扫盘，跳过本次扫盘")
			continue
		}

		fmt.Println("已开始新一阶段扫盘")
		mu.Lock()
		isLock := true
		mu.Unlock()
		go scanDir(roots, &isLock)
	}
}
