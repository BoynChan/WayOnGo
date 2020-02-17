package pool

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/**
模拟数据库的连接池

author:Boyn
date:2020/2/17
*/
const (
	maxGoroutines   = 25 // 最大并发数量
	pooledResources = 2  // 池中资源的数量
)

type dbConnection struct {
	ID int32
}

func (dbConnection *dbConnection) Close() error {
	fmt.Println("Close: Connection ", dbConnection.ID)
	return nil
}

// idCounter用于给每个连接分配一个独一无二的ID
var idCounter int32

// createConnection是一个工厂函数
// 调用这个函数,会返回一个新连接
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	fmt.Println("Create: Connection", id)
	return &dbConnection{id}, nil
}

func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := New(createConnection, pooledResources)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//使用连接池中的资源进行查询
	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	fmt.Println("Done")
	p.Close()
}

// 用于测试连接池
func performQueries(query int, p *Pool) {
	conn, err := p.Acquire()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("QID[%d] CID[%d] \n", query, conn.(*dbConnection).ID)
}
