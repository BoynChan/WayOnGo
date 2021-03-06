package memo

import "errors"

/**
并发的非阻塞缓存
author:Boyn
date:2020/2/18
*/
type entry struct {
	res   result
	ready chan int // 当结果准备好时,会将通道关闭.
}

type request struct {
	key      string
	response chan<- result
}

type Memo struct {
	requests chan request
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

var ErrStopGet = errors.New("function stopped")

var cache = make(map[string]*entry)

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) TryGet(key string, done chan interface{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	select {
	case res := <-response:
		return res.value, res.err
	case <-done:
		delete(cache, key)
		return nil, ErrStopGet
	}
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{
				ready: make(chan int),
			}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
