package list

import "fmt"

type LinkedList struct {
	head *Node
	size int
}
type Node struct {
	Val  int
	Next *Node
}

func New(arr []int) LinkedList {
	l := LinkedList{}
	if len(arr) < 1 {
		return l
	}
	for _, v := range arr {
		l.Add(v)
	}

	return l
}

func (l *LinkedList) Add(val int) {
	if l.head == nil {
		l.head = &Node{
			Val:  val,
			Next: nil,
		}
		l.size++
		return
	}
	iterator := l.head
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = &Node{
		Val:  val,
		Next: nil,
	}
	l.size++
}

func (l *LinkedList) Remove(val int) {
	if l.head == nil {
		return
	}
	// случай когда это head
	if l.head.Val == val {
		l.head = l.head.Next
		l.size--
		return
	}
	iterator := l.head
	for iterator != nil || iterator.Next != nil {
		if iterator.Next.Val == val {
			iterator.Next = iterator.Next.Next
			l.size--
			return
		}
		iterator = iterator.Next
	}
}

func (l *LinkedList) Print() {
	node := l.head
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Println()
}
