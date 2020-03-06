package main

import (
	"7-days-GdCache/GdCache"
	"fmt"
	"log"
	"net/http"
)

// Author:Boyn
// Date:2020/3/6

var db = map[string]string{
	"Tom":  "630",
	"Jack": "1630",
	"Sam":  "2630",
	"Bob":  "3630",
	"Carl": "4630",
}

func main() {
	GdCache.NewGroup("scores", 2<<10, GdCache.GetterFunc(func(key string) ([]byte, error) {
		fmt.Printf("DB - [取值] - %s\n", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s 不存在", key)
	}))

	addr := "localhost:9999"
	peers := GdCache.NewHttpPool(addr)
	fmt.Println("GdCache正在运行:" + addr)
	log.Fatal(http.ListenAndServe(addr, peers))

}
