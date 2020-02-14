package main

import (
	"io"
	"log"
	"net"
	"os"
)

/**
可以接收与发送消息的TCP客户端
author:Boyn
date:2020/2/14
*/

func main() {
	conn, err := net.Dial("tcp", "localhost:8999")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go messageCopy(os.Stdout, conn)
	messageCopy(conn, os.Stdin)
}

func messageCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
