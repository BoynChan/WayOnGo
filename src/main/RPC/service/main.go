package main

import (
	"log"
	"net"
	"net/rpc"
)

//Author: Boyn
//Date: 2020/4/2

func main() {
	_ = RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":12346")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
