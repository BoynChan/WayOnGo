package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// 一致性哈希算法节点的实现
// Author:Boyn
// Date:2020/3/6

// 将节点的名字 映射到 0 - 2^32-1 中
type Hash func(data []byte) uint32

type Map struct {
	hash     Hash           // 用于产生与定位节点的哈希函数
	replicas int            // 虚拟节点的个数
	keys     []int          // 有序的节点列表,即哈希环
	hashMap  map[int]string // 虚拟节点和真实节点的映射表 键为虚拟节点的哈希值,值为真实节点的名称
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		keys:     make([]int, 0),
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// 传入多个Key,批量增加真实节点
// 每一个key对其做不同的hash,重复添加 replicas 次
// 添加完成后,对数组进行排序
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))

	// 使用二分查找法 找到第一个比hash大的值
	index := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[index%len(m.keys)]]
}
