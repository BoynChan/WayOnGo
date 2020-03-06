package GdCache

import (
	"7-days-GdCache/GdCache/consistenthash"
	"7-days-GdCache/GdCache/pb"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/golang/protobuf/proto"
)

// 提供被其他节点访问的能力(基于http)
// Author:Boyn
// Date:2020/3/6
const defaultBasePath = "/_gdcache/"
const defaultReplicas = 50

type HTTPPool struct {
	self        string // 记录自己的地址,主机名与端口 如boyn.top
	basePath    string // 记录通信地址的前缀,即如果地址为 /_gdcache/ 则为通信地址
	mu          sync.Mutex
	peers       *consistenthash.Map    // 一致性哈希算法的节点
	httpGetters map[string]*httpGetter // 将远程节点与本地变量进行对应
}

func NewHttpPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = consistenthash.New(defaultReplicas, nil)
	p.peers.Add(peers...)
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}

// 根据key来选择一个Peer
func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}

func (p *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[ServeHttp %s] %s\n", p.self, fmt.Sprintf(format, v...))
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

	body, err := proto.Marshal(&pb.Response{Value: view.ByteSlice()})
	if err != nil {
		http.Error(w, fmt.Sprintf("缓存请求出错: %s  %s", key, err.Error()), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(body)
}

type httpGetter struct {
	baseURL string
}

func (h *httpGetter) Get(in *pb.Request, out *pb.Response) error {
	u := fmt.Sprintf("%v%v/%v", h.baseURL, url.QueryEscape(in.Group), url.QueryEscape(in.Key))
	res, err := http.Get(u)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("请求状态出错 url:%s status_code %d \n", u, res.StatusCode)
	}
	if res.Header.Get("Content-Type") != "application/octet-stream" {
		return fmt.Errorf("请求头部出错 url:%s Content-Type:%s\n", u, res.Header.Get("Content-Type"))
	}
	bytes, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return fmt.Errorf("解析出错 url:%s err:%s\n", u, err.Error())
	}
	if err = proto.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("解析出错 url:%s err:%s\n", u, err.Error())
	}
	return nil
}
