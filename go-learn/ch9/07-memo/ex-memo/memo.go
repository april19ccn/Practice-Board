/*
互斥锁版本：

更传统的共享内存并发模型

显式管理临界区

计算时释放锁以提高并发性
*/
package memo

import (
	"fmt"
	"sync"
)

// Func is the type of the function to memoize.
type Func func(key string, cancel chan struct{}) (interface{}, error)

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

func (memo *Memo) Get(key string, cancel chan struct{}) (value interface{}, err error) {
	memo.Lock()
	fmt.Println("-----------")
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.Unlock() // 释放锁，允许其他 goroutine 访问缓存（但会阻塞在 <-e.ready）

		// 计算函数结果（耗时操作，无锁状态下执行）
		e.res.value, e.res.err = memo.f(key, cancel)
		fmt.Println(key + "*****")
		select {
		case <-cancel:
			fmt.Println(key + "计算阶段被取消")
			memo.Lock()
			delete(memo.cache, key)
			memo.Unlock()
			return nil, fmt.Errorf("operation cancelled -c")
		default:
			fmt.Println(key + "zzzz")
			close(e.ready) // 广播结果已就绪（所有等待的 goroutine 被唤醒）
		}
	} else {
		memo.Unlock() // 释放锁，仅需等待结果就绪
		<-e.ready     // 阻塞直到 e.ready 关闭（说明结果已计算完成）

		select {
		case <-cancel:
			fmt.Println(key + "等待阶段被取消")
			return nil, fmt.Errorf("operation cancelled -w")
		case <-e.ready:
		}
	}
	fmt.Println(key + "xxxx")
	return e.res.value, e.res.err
}
