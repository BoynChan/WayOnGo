package clock2

import (
	"io"
	"log"
	"net"
	"time"
)

/**
Clock2是一个支持多个客户端进行连接的服务器
其代码与Clock1十分相似,但是在处理连接的函数中
使用go进行修饰,使得它可以每次调用都进入一个独立的goroutine
author:Boyn
date:2020/2/14
*/

func Listen() {
	listener, err := net.Listen("tcp", "localhost:8999")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(connection net.Conn) {
	defer connection.Close()
	for {
		_, err := io.WriteString(connection, time.Now().Format("15:04:05\n"))
		if err != nil {
			//该语句在客户端断开连接时发生
			return
		}
		time.Sleep(1 * time.Second)
	}
}
