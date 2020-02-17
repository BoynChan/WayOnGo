package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

/**
runner包用于展示如何通过通道来获取程序的运行时间
如果运行时间太长,还可以使用runner包来终止程序运行
runner模式可以作为定时任务或监控任务进行

author:Boyn
date:2020/2/17
*/

type Runner struct {
	// interrupt通道用于报告从操作系统中发送的信号
	interrupt chan os.Signal

	// complete通道报告处理任务已经完成
	complete chan error

	// timeout通道报告处理任务已超时
	timeout <-chan time.Time

	// tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// 这个错误信息会在任务超时的时候返回
var ErrTimeout = errors.New("received timeout")

// 这个错误信息会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New返回一个新的,准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		// 在这里使用了time包内置的定时器,经过时间d后,向通道发送当前时间
		// 在程序中,一旦timeout接收到信息,说明程序超时,将程序结束
		timeout: time.After(d),
		tasks:   nil,
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	// 接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	//当任务处理完成的时候发出的信号
	case err := <-r.complete:
		return err
	//当任务超时的时候发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// run执行每一个已经注册了的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {

		// 检测操作系统中的中断讯号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已经注册的任务
		task(id)
	}
	return nil
}

// 检测是否接收到中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	//如果接收到了,即进行打断
	case <-r.interrupt:

		//停止接收后序信号
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
