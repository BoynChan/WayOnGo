package GdCache

// Author:Boyn
// Date:2020/3/6

// 根据传入的key 在一致性哈希选择节点中选择对应的节点(PeerGetter)
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// 从对应group中查找缓存值
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
