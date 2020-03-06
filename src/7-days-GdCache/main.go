package main

import (
	"7-days-GdCache/GdCache"
	"flag"
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

func createGroup() *GdCache.Group {
	return GdCache.NewGroup("scores", 2<<10, GdCache.GetterFunc(func(key string) ([]byte, error) {
		fmt.Printf("DB - [取值] - %s\n", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s 不存在", key)
	}))

}

// 用于启动缓存服务器
// 创建HTTP Pool 添加节点信息,注册到gd中
// 并启动HTTP服务
func startCacheServer(addr string, addrs []string, gd *GdCache.Group) {
	peers := GdCache.NewHttpPool(addr)
	peers.Set(addrs...)
	gd.RegisterPeers(peers)
	fmt.Println("GdCache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

// 启动API服务,与用户进行感知
func startAPIServer(apiAddr string, gd *GdCache.Group) {
	http.Handle("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		view, err := gd.Get(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = w.Write(view.ByteSlice())
	}))
	fmt.Printf("API服务器在 %s 运行\n", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "Geecache server port")
	flag.BoolVar(&api, "api", true, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		18001: "http://localhost:18001",
		18002: "http://localhost:18002",
		18003: "http://localhost:18003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}
	gd := createGroup()

	if api {
		go startAPIServer(apiAddr, gd)
	}
	startCacheServer(addrMap[port], addrs, gd)

}
