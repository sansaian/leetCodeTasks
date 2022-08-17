package list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(arr []int) *ListNode {

	if len(arr) == 0 {
		return nil
	}
	l := ListNode{Val: arr[0]}
	for _, v := range arr[1:] {
		l.Add(v)
	}

	return &l
}

func (l *ListNode) Add(val int) {
	if l == nil {
		l.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		return
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (l *ListNode) Remove(val int) {
	if l == nil {
		return
	}
	if l.Val == val {
		l = l.Next
		return
	}
	iterator := l
	for iterator.Next != nil {
		if iterator.Next.Val == val {
			iterator.Next = iterator.Next.Next
			return
		}
		iterator = iterator.Next
	}
}

func (l *ListNode) Print() {
	node := l
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Println()
}
