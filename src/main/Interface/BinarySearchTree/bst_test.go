package BinarySearchTree

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	b := NewBST()
	b.Insert(1)
	b.Insert(2)
	b.Insert(3)
	b.Insert(4)
	b.Insert(5)
	b.PrintTree()
	fmt.Println(b.Depth())
}
