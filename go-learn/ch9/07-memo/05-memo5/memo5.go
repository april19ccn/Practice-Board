/*
通道版本：(基于CSP模型：使用通道而不是互斥锁来协调并发访问)

更符合Go的"不要通过共享内存来通信，而应该通过通信来共享内存"哲学

使用通道协调所有操作

监控goroutine作为唯一有权访问缓存的主体
*/
package memo5

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type request struct {
	key      string
	response chan<- result
}

type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f) // 启动监控goroutine
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response} // 发送请求
	res := <-response                       // 等待响应
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests { // 循环处理请求
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // 异步调用函数
		}
		go e.deliver(req.response) // 异步传递结果
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready) // 广播结果就绪
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready         // 等待结果就绪
	response <- e.res // 发送结果
}
