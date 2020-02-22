package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Author:Boyn
// Date:2020/2/22

// 基于HTTP协议传输的RPC客户端

// 用于传送相乘的参数
type Args struct {
	A, B int
}

// 用于接收相除的结果,分别为商和余数
type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:8019")
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiple", args, &reply)
	if err != nil {
		log.Fatal("调用失败:", err)
	}
	fmt.Printf("Arith:%d*%d=%d\n", args.A, args.B, reply)
	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("调用失败:", err)
	}
	fmt.Printf("Arith:%d➗%d=%d...%d\n", args.A, args.B, quot.Quo, quot.Rem)

}
