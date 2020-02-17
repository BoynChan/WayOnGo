package work

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
author:Boyn
date:2020/2/17
*/
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	fmt.Println(m.name)
	time.Sleep(time.Second)
}

// 在maxGoroutines=2的情况下,名字是两个两个蹦出来的
func TestWork(t *testing.T) {
	p := New(2)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			//创建一个namePrinter并提供指定的名字
			np := namePrinter{name: name}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
