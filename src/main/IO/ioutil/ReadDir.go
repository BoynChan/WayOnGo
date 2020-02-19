package ioutil

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Author:Boyn
// Date:2020/2/19

// 这个函数将列出给定的路径下所有文件
func listAll(path string, curHire int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curHire; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			listAll(fmt.Sprintf("%s%s%s", path, string(os.PathSeparator), info.Name()), curHire+1)
		} else {
			for tmpHier := curHire; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
