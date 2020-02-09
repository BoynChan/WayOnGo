package BinarySearchTree

import (
	"fmt"
	"math"
)

type BST struct {
	Tree
}

func NewBST() *BST {
	return &BST{}
}

func (B *BST) Insert(val int) {
	if B.root == nil {
		B.root = NewTreeNode(val)
		return
	}
	preNode := B.root
	node := B.root
	for node != nil {
		preNode = node
		if node.val > val {
			node = node.left
		} else {
			node = node.right
		}
	}
	if preNode.val > val {
		preNode.left = NewTreeNode(val)
	} else {
		preNode.right = NewTreeNode(val)
	}
}

func (B BST) PrintTree() {
	Print(B.root, 0)
}

func Print(node *TreeNode, depth int) {
	if node == nil {
		return
	}
	for i := depth; i > 0; i-- {
		fmt.Printf(" ")
	}
	fmt.Println(node.val)
	Print(node.left, depth+1)
	Print(node.right, depth+1)
}

func (B BST) Depth() int {
	return Depth(B.root)
}

func Depth(node *TreeNode) int {
	if node == nil {
		return 0
	} else {
		return int(math.Max(float64(Depth(node.left)), float64(Depth(node.right))) + 1)
	}
}

func (B BST) LeafCount() int {
	panic("implement me")
}

func (B BST) PreOrder() {

}

func (B BST) InOrder() {
	panic("implement me")
}

func (B BST) PostOrder() {
	panic("implement me")
}
