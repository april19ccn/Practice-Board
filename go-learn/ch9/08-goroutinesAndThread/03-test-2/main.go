package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 计算密集型任务：计算斐波那契数列
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 执行计算密集型任务的worker
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		results <- fibonacci(n)
	}
}

func main() {
	fmt.Println("计算密集型并发程序与GOMAXPROCS性能测试")
	fmt.Printf("CPU核心数量: %d\n", runtime.NumCPU())

	// 测试不同的GOMAXPROCS值
	maxProcsValues := []int{1, 2, 4, 8, 16, 32, 64}
	numJobs := 40
	fibNumber := 40 // 计算斐波那契数列的第40个数，这是一个计算密集型任务

	for _, maxProcs := range maxProcsValues {
		runtime.GOMAXPROCS(maxProcs)

		// 创建任务和结果channel
		jobs := make(chan int, numJobs)
		results := make(chan int, numJobs)

		// 启动worker
		var wg sync.WaitGroup
		start := time.Now()

		// 创建与GOMAXPROCS相同数量的worker
		for w := 1; w <= maxProcs; w++ {
			wg.Add(1)
			go worker(w, jobs, results, &wg)
		}

		// 发送任务
		for j := 1; j <= numJobs; j++ {
			jobs <- fibNumber
		}
		close(jobs)

		// 等待所有worker完成
		wg.Wait()
		close(results)

		elapsed := time.Since(start)

		// 收集结果（虽然我们不使用，但需要从channel中读取）
		var fibResults []int
		for result := range results {
			fibResults = append(fibResults, result)
		}

		fmt.Printf("GOMAXPROCS=%d, 耗时: %v, 处理了 %d 个任务\n",
			maxProcs, elapsed, len(fibResults))
	}

	// 恢复默认设置
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("\n测试完成，GOMAXPROCS已恢复为默认值: %d\n", runtime.GOMAXPROCS(0))
}
