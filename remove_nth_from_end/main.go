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

func (l *ListNode) print() {
	node := l
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Println()
}

func main() {
	list := newListNode()
	list.print()
	removeNthFromEnd(list, 3)
	list.print()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	startPtr := head
	endPtr := head
	size := 1
	for i := 0; i < n; i++ {
		if endPtr.Next != nil {
			endPtr = endPtr.Next
			size++
		}
	}
	for endPtr.Next != nil {
		startPtr = startPtr.Next
		endPtr = endPtr.Next
		size++
	}
	if size == n {
		return head.Next
	}
	removeElement := startPtr.Next
	startPtr.Next = removeElement.Next
	return head
}
