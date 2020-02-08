package main

/*
单链表
*/
import "fmt"

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
	size int
}

func NewList() *List {
	return &List{
		head: &Node{
			val:  -1,
			next: nil,
		},
		size: 0,
	}
}

func (l *List) AddNode(val int) {
	node := l.head
	for node.next != nil {
		node = node.next
	}
	newNode := &Node{
		val: val,
	}
	l.size++
	node.next = newNode
}

func (l *List) Reverse() {
	node := l.head.next
	for node != nil {
		fmt.Printf("%d ", node.val)
		node = node.next
	}
	fmt.Println()
}

func main() {
	list := NewList()
	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)
	list.Reverse()
}
