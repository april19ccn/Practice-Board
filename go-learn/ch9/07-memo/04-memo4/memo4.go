/*
互斥锁版本：

更传统的共享内存并发模型

显式管理临界区

计算时释放锁以提高并发性
*/
package memo4

import "sync"

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

// A Memo caches the results of calling a Func.
type Memo struct {
	sync.Mutex
	f     Func
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.Unlock() // 释放锁，允许其他 goroutine 访问缓存（但会阻塞在 <-e.ready）

		// 计算函数结果（耗时操作，无锁状态下执行）
		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // 广播结果已就绪（所有等待的 goroutine 被唤醒）
	} else {
		memo.Unlock() // 释放锁，仅需等待结果就绪
		<-e.ready     // 阻塞直到 e.ready 关闭（说明结果已计算完成）
	}
	return e.res.value, e.res.err
}

// 并发执行示例
// 假设两个 goroutine（G1 和 G2）同时请求同一个 key：

// 1 -- G1 先获取锁：

// 发现 cache[key] 为空，创建 entry 并释放锁。

// 开始计算 memo.f(key)。

// 2 -- G2 获取锁：

// 发现 cache[key] 已存在（G1 创建的 entry），释放锁。

// 执行 <-e.ready 阻塞，等待结果。

// 3 -- G1 计算完成：

// 关闭 e.ready，唤醒所有等待的 goroutine（包括 G2）。

// 4 -- G1 和 G2 均返回结果。
