package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode() *ListNode {

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

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var last, tmp *ListNode

	for head.Next != nil {
		tmp = head.Next
		head.Next, last = last, head
		head = tmp
	}
	head.Next = last
	return head
}

func main() {
	head := newListNode()
	head.Print()
	reverseList(head)
	head.Print()
}
