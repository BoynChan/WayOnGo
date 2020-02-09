package BinarySearchTree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val: val}
}

//操作器,可以对树进行输出,深度计算和叶子结点个数的统计
type Operate interface {
	Insert(int)
	PrintTree()
	Depth() int
	LeafCount() int
}

//遍历器,对树进行前序,中序和后序遍历
type Order interface {
	PreOrder()
	InOrder()
	PostOrder()
}

//普通树的定义
type Tree struct {
	root *TreeNode
}
