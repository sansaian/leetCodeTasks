package main

type List struct {
	root ListNode // sentinel list element, only &root, root.prev, and root.next are used
	len  int      // current list length excluding (this) sentinel element
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if node.Next == nil {
		node = nil
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {

}
