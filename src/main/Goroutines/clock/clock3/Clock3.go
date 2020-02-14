package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/**
支持使用-port参数来指定端口,并使用-timezone参数来指定时区,如中国为+8区
author:Boyn
date:2020/2/14
*/

func Listen(port int, timezone string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "Listening on port:%d\nTimeZone:%s\n", port, timezone)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn, timezone)
	}
}

func handleConn(connection net.Conn, timezone string) {
	defer connection.Close()
	location, _ := time.LoadLocation(timezone)
	for {
		_, err := io.WriteString(connection, time.Now().In(location).Format("15:04:05"))
		if err != nil {
			//该语句在客户端断开连接时发生
			return
		}
		time.Sleep(1 * time.Second)
	}
}

var port = flag.Int("port", 8999, "-port:8999")
var timezone = flag.String("timezone", "Asia/Chongqing", "-timezone:Asia/Chongqing  # represent china")

/*func main(){
	flag.Parse()
	Listen(*port,*timezone)
}*/
