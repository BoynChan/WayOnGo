package GdCache

import (
	lru "7-days-GdCache/GdCache/lru"
	"sync"
)

// Author:Boyn
// Date:2020/3/5

type cache struct {
	mu         sync.Mutex // 互斥锁
	lru        *lru.Cache // lru缓存
	cacheBytes int64      // 已经缓存的字节
}

// 将元素添加到缓存中
func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lruLazyInit()
	}
	c.lru.Add(key, value)
}

func (c *cache) lruLazyInit() {
	c.lru = lru.New(c.cacheBytes, nil)
}

// 获取缓存的值
func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lruLazyInit()
	}
	if value, ok := c.lru.Get(key); ok {
		return value.(ByteView), ok
	}
	return
}
