package GdCache

import (
	"fmt"
	"sync"
)

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

// load函数在后面会进行拓展
// 目前的功能是从本地(也就是我们指定的函数中)进行佳在
func (g *Group) load(key string) (value ByteView, err error) {
	return g.getLocally(key)
}

// 从本地(即开始指定的函数中)进行加载
func (g *Group) getLocally(key string) (ByteView, error) {
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
