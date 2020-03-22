package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Author:Boyn
// Date:2020/3/22

func main() {
	dir := "F:\\Code\\python\\大创\\数据集\\广场"
	files, _ := ioutil.ReadDir(dir)
	counter := 0
	for _, v := range files {
		if v.IsDir() {
			continue
		}

		err := os.Rename(dir+"\\"+v.Name(), dir+"\\"+strconv.Itoa(counter)+".jpg")
		counter++
		if err != nil {
			fmt.Println(err)
		}
	}
}
