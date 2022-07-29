package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewGenerateList(n int) *ListNode {

	var last *ListNode
	for i := 5; i > 0; i-- {
		node := &ListNode{
			Val:  i * n,
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	val := 0
	iterator2 := list2
	iterator1 := list1
	switch {
	case list1 == nil:
		return list2
	case list2 == nil:
		return list1
	case list1.Val > list2.Val:
		iterator2 = list2.Next
		val = list2.Val
	case list1.Val <= list2.Val:
		iterator1 = list1.Next
		val = list1.Val
	}

	resultList := &ListNode{
		Val:  val,
		Next: nil,
	}
	head := resultList

	for iterator1 != nil && iterator2 != nil {
		switch {
		case iterator1.Val <= iterator2.Val:
			val = iterator1.Val
			iterator1 = iterator1.Next
		case iterator1.Val > iterator2.Val:
			val = iterator2.Val
			iterator2 = iterator2.Next
		}
		resultList.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		resultList = resultList.Next
	}
	if iterator1 == nil {
		resultList.Next = iterator2
		return head
	}
	resultList.Next = iterator1
	return head
}

func main() {
	list1 := NewGenerateList(1)
	list1.Print()
	list2 := NewGenerateList(3)
	list2.Print()
	result := mergeTwoLists(list1, list2)
	result.Print()
}
