package bank

import (
	"sync"
)

/**
使用读写锁分离读取操作和写入操作
对于读写锁来说,多个读操作是兼容的,而写操作与另外的读操作和另外的写操作都是不兼容的
author:Boyn
date:2020/2/18
*/

var (
	rw             sync.RWMutex
	balanceRWMutex int
)

func DepositByRWMutex(amount int) {
	rw.Lock()
	defer rw.Unlock()
	balanceMutex += amount
}

func WithdrawByRWMutex(amount int) bool {
	rw.Lock()
	defer rw.Unlock()
	if balanceMutex > amount {
		balanceMutex -= amount
		return true
	} else {
		return false
	}
}

func BalanceByRWMutex() int {
	//在此处,读操作是用的读锁,前面都是用的写锁
	rw.RLock()
	defer rw.RUnlock()
	return balanceMutex
}
