package semaphore

/**
用于控制Goroutines并发量的信号量工具
author:Boyn
date:2020/2/15
*/

type semaphore struct {
	n       int // 并发数量
	current int // 当前可用数
	tokens  chan int
}

/*
semaphore本身的数据结构不导出,只暴露API
*/
func NewSemaphore(n int) *semaphore {
	return &semaphore{n: n, current: n, tokens: make(chan int, n)}
}

/*
获取可用
*/
func (s *semaphore) Acquire() {
	s.tokens <- 1
	s.current--
}

/*
释放可用
*/
func (s *semaphore) Release() {
	<-s.tokens
	s.current++
}

/*
查看当前可用
*/
func (s *semaphore) Current() int {
	return s.current
}
