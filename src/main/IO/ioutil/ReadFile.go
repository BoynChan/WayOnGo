package ioutil

import "io/ioutil"

// Author:Boyn
// Date:2020/2/19

// io util的读取文件操作
func readFile(filename string) (file []byte, err error) {
	file, err = ioutil.ReadFile(filename)
	return
}
