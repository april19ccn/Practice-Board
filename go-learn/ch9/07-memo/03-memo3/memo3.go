// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 272.

//!+

// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo3

import "sync"

// A Memo caches the results of calling a Func.
type Memo struct {
	sync.Mutex
	f     Func
	cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

/*
减少锁竞争：多个goroutine可以同时执行f(key)操作，而不需要相互等待。这对于处理不同key的请求特别有效，因为每个key的计算可以并行进行。

提高并发性：当一个goroutine在计算f(key)时，其他goroutine可以继续检查缓存或计算其他key，从而充分利用CPU和I/O资源。

缩短临界区：锁的持有时间大大缩短，减少了goroutine阻塞的时间，提高了系统的整体吞吐量。
*/
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.Lock()
	res, ok := memo.cache[key]
	memo.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Between the two critical sections, several goroutines
		// may race to compute f(key) and update the map.
		memo.Lock()
		memo.cache[key] = res
		memo.Unlock()
	}
	return res.value, res.err
}
