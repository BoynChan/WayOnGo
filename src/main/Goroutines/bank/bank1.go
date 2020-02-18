package bank

/**
无原子性保证的银行转账程序
author:Boyn
date:2020/2/18
*/

// 储蓄
var balance int

// 将钱放入存款
func DepositUnsafe(amount int) {
	balance += amount
}

func WithdrawUnsafe(amount int) {
	balance -= amount
}

// 返回储蓄数目
func BalanceUnsafe() int {
	return balance
}
