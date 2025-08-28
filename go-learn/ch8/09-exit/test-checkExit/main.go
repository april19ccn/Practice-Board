// 本测试是为了理解这段话
// 现在当取消发生时，所有后台的goroutine都会迅速停止并且主函数会返回。
// 当然，当主函数返回时，一个程序会退出，而我们又无法在主函数退出的时候确认其已经释放了所有的资源（译注：因为程序都退出了，你的代码都没法执行了）。
// 这里有一个方便的窍门我们可以一用：取代掉直接从主函数返回，我们调用一个panic，然后runtime会把每一个goroutine的栈dump下来。
// 如果main goroutine是唯一一个剩下的goroutine的话，他会清理掉自己的一切资源。
// 但是如果还有其它的goroutine没有退出，他们可能没办法被正确地取消掉，也有可能被取消但是取消操作会很花时间；
// 所以这里的一个调研还是很有必要的。我们用panic来获取到足够的信息来验证我们上面的判断，看看最终到底是什么样的情况。
package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d 正在退出...\n", id)
			return
		default:
			// 模拟工作负载
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func leakyWorker(ctx context.Context, id int) {
	// 添加启动通知
	fmt.Printf("Leaky worker %d 启动...\n", id)

	// 这个worker没有检查ctx.Done()，会导致泄露
	for {
		fmt.Printf("Leaky worker %d 正在工作...\n", id)
		time.Sleep(200 * time.Millisecond)
	}
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

	// 设置GOMAXPROCS确保多核运行
	runtime.GOMAXPROCS(runtime.NumCPU())

	ctx, cancel := context.WithCancel(context.Background())

	// 启动正常工作的worker
	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	// 启动有泄露问题的worker
	for i := 1; i <= 2; i++ {
		go leakyWorker(ctx, i)
	}

	// 确保goroutine有时间启动
	fmt.Println("等待goroutine启动...")
	time.Sleep(500 * time.Millisecond)

	// 打印当前goroutine数量
	fmt.Printf("当前goroutine数量: %d\n", runtime.NumGoroutine())

	// 让程序运行一段时间
	fmt.Println("程序运行中...")
	time.Sleep(1 * time.Second)

	// 发送取消信号
	fmt.Println("发送取消信号...")
	cancel()

	// 等待一段时间让正常worker退出
	fmt.Println("等待正常worker退出...")
	time.Sleep(500 * time.Millisecond)

	// 再次打印goroutine数量
	fmt.Printf("取消后goroutine数量: %d\n", runtime.NumGoroutine())

	// 使用panic来检查goroutine状态
	fmt.Println("使用panic检查goroutine状态...")
	panic("调试：检查goroutine退出状态")
}

// 解决方案：生产环境的正确做法：使用 sync.WaitGroup 来协同等待所有 worker goroutine 安全退出。

// 程序结束了，其他协程不都会被销毁，为什么担心资源没有释放？
