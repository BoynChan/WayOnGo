package bank

import "sync"

/**
author:Boyn
date:2020/2/18
*/

var (
	mu           sync.Mutex
	balanceMutex int
)

func DepositByMutex(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balanceMutex += amount
}

func WithdrawByMutex(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	if balanceMutex > amount {
		balanceMutex -= amount
		return true
	} else {
		return false
	}
}

func BalanceByMutex() int {
	mu.Lock()
	defer mu.Unlock()
	return balanceMutex
}
