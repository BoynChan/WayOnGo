package bank

/**
在Go中,通道是一种很好的避免数据竞争的方式
因为通道本身就是线程安全的
author:Boyn
date:2020/2/18
*/

var deposits = make(chan int)          // 存款通道
var withdraw = make(chan withdrawNode) // 存款通道
var balances = make(chan int)          // 用于放置现在数量的存款的通道

func DepositByChannel(amount int) {
	deposits <- amount
}

func BalanceByChannel() int {
	return <-balances
}

//使用一个结构体来表示取款的额度以及返回结果
type withdrawNode struct {
	result chan bool
	amount int
}

func WithdrawByChannel(amount int) bool {
	ch := make(chan bool)
	w := withdrawNode{
		result: ch,
		amount: amount,
	}
	withdraw <- w
	return <-ch
}

func teller() {
	var balance int // 实际存储存款的变量被限制在一个函数中,并多个协程对其的读写操作用select来扁平化称为线程安全的
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdraw:
			if w.amount > balance {
				w.result <- false
			} else {
				w.result <- true
				balance -= w.amount
			}
		case balances <- balance:
		}
	}
}

/**
放在init中,启动teller
*/
func init() {
	go teller()
}
