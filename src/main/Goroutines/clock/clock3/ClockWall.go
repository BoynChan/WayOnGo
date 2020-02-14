package main

import (
	"fmt"
	"net"
	"os"
	"text/tabwriter"
)

/**
一个连接多个时钟服务器的客户端,以表格的方式打印
author:Boyn
date:2020/2/14
*/

func main() {
	Chongqing, _ := net.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", 8990))
	Tokyo, _ := net.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", 8991))
	Canberra, _ := net.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", 8992))
	London, _ := net.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", 8993))
	printTimeTitle()
	buf := make([]byte, 32)
	for {
		Chongqing.Read(buf[0:8])
		Tokyo.Read(buf[8:16])
		Canberra.Read(buf[16:24])
		London.Read(buf[24:32])
		fmt.Printf("    %s    ", string(buf[0:8]))
		fmt.Printf("    %s    ", string(buf[8:16]))
		fmt.Printf("    %s    ", string(buf[16:24]))
		fmt.Printf("    %s    ", string(buf[24:32]))
		fmt.Printf("\r")
	}
}

func printTimeTitle() {
	const format = "\r%v\t%v\t%v\t%v\t"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 16, 3, ' ', 0)
	fmt.Fprintf(tw, format, "Asia/Chongqing", "Asia/Tokyo", "Australia/Canberra", "Europe/London\n")
	fmt.Fprintf(tw, format, "--------------", "----------", "------------------", "-------------\n")
	tw.Flush()
}
