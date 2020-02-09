package logger

import (
	"fmt"
	"os"
)

/*
创建一个控制台类型的日志写入器
*/
type consoleWriter struct {
}

/*
日志写入的实现
将数据格式化后,写入标准输出中
*/
func (c *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

func NewConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
