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

func main() {

}

func hasCycle(head *ListNode) bool {
	for slow, fast := head, head; slow != nil && fast != nil; {
		slow = slow.Next
		fast = fast.Next
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next
		if slow == fast {
			return true
		}
	}
	return false
}
