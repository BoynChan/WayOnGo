package bank

import (
	"fmt"
	"sync"
	"testing"
)

/**
author:Boyn
date:2020/2/18
*/

/*
这个程序在1000次循环后,其结果往往不等于100000
因为产生了竞争条件
无论何时,只要有多个协程同时访问一个变量,就会其中至少有两个是写操作的时候,就会发送数据竞争
*/
func TestBank1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			DepositUnsafe(100)
		}()
	}
	wg.Wait()
	fmt.Printf("BalanceUnsafe=%d\n", BalanceUnsafe())
}

/*
在这个函数中,使用通道来将线程不安全的写操作转为安全的
*/
func TestBank2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1500)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			DepositByChannel(100)
		}()
	}
	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			WithdrawByChannel(100)
		}()
	}
	wg.Wait()
	fmt.Printf("BalanceUnsafe=%d\n", BalanceByChannel())
}

/*
在这个函数中,使用互斥锁来将线程不安全的写操作转为安全的
*/
func TestBank3(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1500)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			DepositByMutex(100)
		}()
	}
	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			WithdrawByMutex(100)
		}()
	}
	wg.Wait()
	fmt.Printf("BalanceUnsafe=%d\n", BalanceByMutex())
}
