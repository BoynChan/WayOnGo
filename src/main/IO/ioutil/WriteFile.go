package ioutil

import (
	"io/ioutil"
	"os"
)

// Author:Boyn
// Date:2020/2/19

// io util的WriteFile操作
// 0666指的是默认的文件写入模式
func WriteFile(filename string, data []byte) {
	_ = ioutil.WriteFile(filename, data, os.FileMode(0666))
}
