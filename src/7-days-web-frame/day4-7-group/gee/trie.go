package gee

import "strings"

// Author:Boyn
// Date:2020/3/3

type node struct {
	pattern  string  // 待匹配路由
	part     string  // 路由中的一部分
	children []*node // 子节点
	isWild   bool    //是否精确匹配
}

// 以当前节点为起点,搜索所有子节点与part相同的节点
// 返回第一个成功匹配的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 以当前节点为起点,搜索所有子节点与part相同的节点
// 返回所有匹配成功的节点,用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 将动态路径插入到前缀树中
// 这里使用递归插入,并将height与被 / 分割的路径数组进行索引的对称
// height的数值就是路径数组的索引值
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 进行前缀树节点的搜索
// 如果已经搜索到了最后一个节点,或者通配( * )节点
// 就返回给上层
// 对每一层来说,会首先将路径数组对应索引处的路径拿出来,并找到该节点中与路径对应的children
// 然后遍历这些子节点进行递归查找
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
