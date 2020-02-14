package clock1

import (
	"io"
	"log"
	"net"
	"time"
)

/**
Clock1是一个简单的TCP服务器,每隔一秒钟,将时间写到与其相连的客户端中
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
		handleConn(conn)
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
