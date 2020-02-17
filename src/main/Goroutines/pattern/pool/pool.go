package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

/**
使用有缓冲的通道实现资源池

author:Boyn
date:2020/2/17
*/

// Pool用于管理在多个Goroutines中共享的资源
// 被管理的资源需要实现io.Closer接口
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// 该错误表示请求了一个已经关闭了的池
var ErrPoolClosed = errors.New("pool has been closed")

// New函数创建一个用于管理资源的池
// 这个池需要一个可以分配新资源的函数,并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}
	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   fn,
	}, nil

}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	//当有资源可用时,会将缓冲通道中的资源取出并返回给上层函数
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
		//当无资源可用时,会新建一个资源并返回
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release将一个使用后的资源放回池中
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	//如果池已经被关闭,那么就将这个资源销毁
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release:", "In queue")
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	// 在清空通道中的资源之前,要先关闭通道
	// 我们在之前说过,关闭了通道之后,只读不写
	// 这可以避免死循环
	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}
