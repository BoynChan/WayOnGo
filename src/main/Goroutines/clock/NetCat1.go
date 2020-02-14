package clock

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

/**
NetCat1是一个简单的,只读的TCP客户端
author:Boyn
date:2020/2/14
*/

func Connect(address string, port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	messageCopy(os.Stdout, conn)
}

func messageCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
