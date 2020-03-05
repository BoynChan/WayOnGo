package lru

import "container/list"

// Author:Boyn
// Date:2020/3/5

type Cache struct {
	maxBytes  int64                         // 允许使用的最大内存
	nbytes    int64                         // 当前已经使用的内存
	ll        *list.List                    // lru中的list
	cache     map[string]*list.Element      // 缓存存放的map
	OnEvicted func(key string, value Value) // 某条记录被移除时的回调函数
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

// 对缓存结构体做初始化,其中可以传入的参数是允许的最大内存
// 以及当剩余内存不足而将缓存元素移除时调用的回调函数
func New(maxBytes int64, onEvicted func(key string, value Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		nbytes:    0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 查找功能的实现
// 每一个list.Element中都是一个entry的指针
// Cache中的map和链表是具有一致性的,确保map中存在的元素,链表中一定存在
// 所以我们可以通过key快速地从map中查找到element元素后,将这个元素移到链表的头部
// 然后将其返回给上层
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 移除最少被访问的节点
// 首先取出队尾的元素,如果取出的元素不为空,就将其从链表中移除
// 然后在map中将这个节点对应的key value删除
// 接着将当前内存占用减去key的长度和value的长度
// 最后执行删除策略的回调函数
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 向缓存中添加元素,如果已存在则更新,不存在则插入.
func (c *Cache) Add(key string, value Value) {
	// 如果元素存在于map中
	// 则仅对元素进行更新
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
		return
	}

	// 如果元素不存在,则创建一个entry将其插入到链表中
	// 然后将其放入map并更新已用内存
	ele := c.ll.PushFront(&entry{
		key:   key,
		value: value,
	})
	c.cache[key] = ele
	c.nbytes += int64(len(key)) + int64(value.Len())

	// 当当前内存超出了最大内存限制的时候,将最老的元素移除
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// 获取已存的元素
func (c *Cache) Len() int {
	return c.ll.Len()
}

// 查询对应元素是否存在
func (c *Cache) IsExist(key string) bool {
	_, ok := c.cache[key]
	return ok
}
