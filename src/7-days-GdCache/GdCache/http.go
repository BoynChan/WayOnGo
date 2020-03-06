package GdCache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// 提供被其他节点访问的能力(基于http)
// Author:Boyn
// Date:2020/3/6
const defaultBasePath = "/_gdcache/"

type HTTPPool struct {
	self     string // 记录自己的地址,主机名与端口
	basePath string // 记录通信地址的前缀,即如果地址为boyn.top/_gdcache/ 则为通信地址
}

func NewHttpPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (p *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[ServeHttp %s] %s\n", p.basePath, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		p.Log("请求地址错误")
	}
	p.Log("%s %s", r.Method, r.URL.Path)

	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "请求地址格式错误", http.StatusBadRequest)
		p.Log("%s 请求地址错误", r.URL.Path)
		return
	}
	groupName := parts[0]
	key := parts[1]
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, fmt.Sprintf("无此组名: %s", groupName), http.StatusBadRequest)
		p.Log("无此组名: %s", groupName)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(w, fmt.Sprintf("缓存请求出错: %s  %s", key, err.Error()), http.StatusBadRequest)
		p.Log("缓存请求出错: %s  %s", key, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	n, err := w.Write(view.ByteSlice())
	fmt.Println(n)
	if err != nil {
		http.Error(w, fmt.Sprintf("缓存请求出错: key:[%s]  %s", key, err.Error()), http.StatusBadRequest)
		p.Log("缓存请求出错: %s  %s", key, err.Error())
		return
	}
}
