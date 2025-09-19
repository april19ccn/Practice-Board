package main

import (
	"fmt"
	"runtime"
	"time"
)

// 流水线阶段：接收输入，处理，发送到输出
func pipelineStage(in <-chan int, out chan<- int, stageNum int) {
	for value := range in {
		// 简单的处理：增加值并传递
		result := value + 1
		out <- result
	}
	close(out)
}

func main() {
	fmt.Println("Go流水线阶段测试")
	fmt.Printf("当前GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("初始goroutine数量: %d\n", runtime.NumGoroutine())

	// 内存状态
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("初始内存分配: %v KB\n", mem.Alloc/1024)

	// 测试不同数量的阶段
	stageCounts := []int{100, 1000, 10000, 100000, 1000000}

	for _, stages := range stageCounts {
		fmt.Printf("\n测试 %d 个阶段...\n", stages)

		// 创建输入和输出channel
		start := time.Now()
		firstIn := make(chan int)
		prevOut := firstIn

		// 创建指定数量的流水线阶段
		for i := 0; i < stages; i++ {
			stageOut := make(chan int)
			go pipelineStage(prevOut, stageOut, i+1)
			prevOut = stageOut
		}

		// 最后的输出channel
		finalOut := prevOut

		// 发送一个值通过整个流水线
		value := 1
		go func() {
			firstIn <- value
			close(firstIn)
		}()

		// 从最终输出接收结果
		result := <-finalOut

		elapsed := time.Since(start)

		// 获取内存和goroutine信息
		runtime.GC() // 强制垃圾回收以获得更准确的内存统计
		runtime.ReadMemStats(&mem)
		goroutines := runtime.NumGoroutine()

		fmt.Printf("结果: %d (期望: %d)\n", result, value+stages)
		fmt.Printf("通过时间: %v\n", elapsed)
		fmt.Printf("内存使用: %v KB\n", mem.Alloc/1024)
		fmt.Printf("当前goroutine数量: %d\n", goroutines)

		// 如果goroutine数量增长过多，提前结束
		if goroutines > 10000 {
			fmt.Println("警告: goroutine数量过多，停止测试")
			break
		}
	}

	fmt.Println("\n测试完成")
}
