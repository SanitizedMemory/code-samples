package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func (node *Node) Push(value int) {
	if node.Next == nil {
		node.Next = &Node{value, nil}
		return
	}
	node.Next.Push(value)
}

func (node *Node) PopByValue(value int) {
	/* If the user tries to pop the first node,
	 * kill them.
	 */
	if node.Next == nil {
		return
	}
	if node.Next.Value == value {
		node.Next = node.Next.Next
		return
	}
	node.Next.PopByValue(value)
}

func (node *Node) buildString() string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("%d -> %s", node.Value, node.Next.buildString())
}

func (node *Node) Print() {
	fmt.Println(node.buildString())
}

func makeLinkedList(nums ...int) *Node {
	linkedList := Node{nums[0], nil}
	head := &linkedList
	for i := 1; i < len(nums); i++ {
		head.Next = &Node{nums[i], nil}
		head = head.Next
	}
	return &linkedList
}

func main() {
	linkedList := *makeLinkedList(1, 2, 3)
	linkedList.Push(40)
	linkedList.PopByValue(2)
	linkedList.Print()

}
