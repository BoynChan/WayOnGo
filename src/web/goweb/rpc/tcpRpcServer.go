package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

// Author:Boyn
// Date:2020/2/22

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

// 对于RPC来说,要用 func (type *Type) funcName(args Args, reply *Reply) error 这种类型的函数才可以被识别
func (t *Arith) Multiple(args Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args Args, reply *Quotient) error {
	if args.B == 0 {
		return errors.New("除数不能为0")
	}
	reply.Quo = args.A / args.B
	reply.Rem = args.A % args.B
	return nil
}

func main() {
	// 这里要是指针
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8019")
	if err != nil {
		fmt.Println(err.Error())
	}
	tcp, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		conn, err := tcp.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
