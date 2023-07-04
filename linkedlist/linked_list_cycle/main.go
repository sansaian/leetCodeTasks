package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewGenerateList() *ListNode {
	var last *ListNode
	for i := 5; i > 0; i-- {
		node := &ListNode{
			Val:  i,
			Next: last,
		}
		last = node
	}
	return last
}

func (l *ListNode) Print() {
	node := l
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Println()
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
