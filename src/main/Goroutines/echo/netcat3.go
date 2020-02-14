package main

import (
	"io"
	"log"
	"net"
	"os"
)

/**
使用channel来同步客户端的输入输出
author:Boyn
date:2020/2/14
*/

func main() {
	conn, err := net.Dial("tcp", "localhost:8999")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan int)
	defer conn.Close()

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("Done")
		done <- 1 //作为信号,发送给main
	}()
	messageCopy(conn, os.Stdin)
	<-done // 接收到协程的结束信号
}
