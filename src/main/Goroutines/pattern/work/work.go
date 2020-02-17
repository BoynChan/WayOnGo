package work

import "sync"

/**
使用无缓冲的通道,实现一个简单的协程池,使其可以接收任务并运行
author:Boyn
date:2020/2/17
*/

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建一个新的工作池
// 其中,在开始工作时已经创建了maxGoroutines个协程用于接收任务并执行
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
