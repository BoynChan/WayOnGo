package memo

import "sync"

/**
并发的非阻塞缓存
author:Boyn
date:2020/2/18
*/
type entry struct {
	res   result
	ready chan int // 当结果准备好时,会将通道关闭.
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]

	//当没有缓存的时候
	if e == nil {
		// 在这里使用了延迟处理的思想,先用entry作为占位,但此时f函数运行还没有出结果
		// 但是cache[key]处已经有赋值了.然后可以解锁并运行函数
		// 当后面有其他协程要进入时,可以获取到e的值不为空,但是会阻塞在<-e.ready处
		// 当请求完毕后,会关闭ready通道,使所有请求通道输出的协程停止阻塞
		e = &entry{
			ready: make(chan int),
		}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // 向所有在这个key处等待结果阻塞的协程发出信号,使其能获取结果
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}
