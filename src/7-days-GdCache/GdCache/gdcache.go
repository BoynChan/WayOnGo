package GdCache

import (
	"fmt"
	"sync"
)

// 外部访问的类,要经过命名空间,才可以到达缓存中
// Author:Boyn
// Date:2020/3/5

// Getter接口是当缓存从本地和网络中都获取不到时,调用用户实现的Getter接口中的回调函数进行获取
type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// Group是缓存的命名空间
// 每一个Group拥有自己的唯一名称,缓存不同属性的值
type Group struct {
	name      string // 命名空间
	getter    Getter // 缓存未命中时从此处拿取值
	mainCache cache  // 实际缓存存储的地方
	peers     PeerPicker
}

func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		fmt.Println("无法注册两次")
		return
	}
	g.peers = peers
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

// 新建一个组,getter必须传入不能为nil
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("Getter is none")
	}
	mu.Lock()
	defer mu.Unlock()
	group := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = group
	return group
}

// 以namespace提取一个组
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

// 从缓存中提取一个值
// 如果缓存中没有这个值,就从传入的函数中进行获取
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.get(key); ok {
		fmt.Printf("GdCache - [缓存命中] - %s\n", key)
		return v, nil
	}

	return g.load(key)
}

// 目前的功能是先在已经注册的API服务器中进行加载
// 如果没有已注册的API服务器,或获取失败,则从本地加载
func (g *Group) load(key string) (value ByteView, err error) {
	if g.peers != nil {
		if peer, ok := g.peers.PickPeer(key); ok {
			if value, err := g.getFromPeer(peer, key); err == nil {
				return value, nil
			}
			fmt.Printf("GdCache - [无法从peer获取key] - key %s\n", key)
		}
	}
	return g.getFromLocal(key)
}

func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	bytes, err := peer.Get(g.name, key)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: bytes}, nil
}

// 从本地(即开始指定的函数中)进行加载
func (g *Group) getFromLocal(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

// 将键值加入到缓存中
func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
