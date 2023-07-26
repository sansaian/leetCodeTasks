package main

import "fmt"

func main() {
	mylist := Constructor()
	mylist.AddAtHead(1)
	mylist.AddAtTail(3)
	mylist.AddAtIndex(1, 2)
	fmt.Println(mylist.Get(1))
	//fmt.Println(mylist.Get(2))
	//fmt.Println(mylist.Get(3))
	//fmt.Println(mylist.Get(4))

	//mylist.AddAtIndex(0, 7)
	//mylist.AddAtIndex(2, -1)
	//mylist.AddAtIndex(4, -1)
	//mylist.DeleteAtIndex(2)
	//mylist.DeleteAtIndex(4)
	fmt.Println()

}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	Next *Node
	Val  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		tail: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	curr := this.head
	for i := 0; curr.Next != nil && i < index; i++ {
		curr = curr.Next
	}
	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Next: this.head, Val: val}
	if this.tail == nil {
		this.tail = node
	}
	this.head = node
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.tail == nil {
		this.AddAtHead(val)
		return
	}
	node := &Node{Next: nil, Val: val}
	this.tail.Next = node
	this.tail = node
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}
	prev := this.head
	newNode := &Node{Next: nil, Val: val}
	for i := 1; i < index; i++ {
		prev = prev.Next
	}
	newNode.Next = prev.Next
	prev.Next = newNode
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}
	this.size--
	if this.size == 0 {
		this.head = nil
		this.tail = nil
	} else if index == 0 {
		this.head = this.head.Next
	} else {
		prev := this.head
		for i := 1; i < index; i++ {
			prev = prev.Next
		}
		if index == this.size {
			this.tail = prev
			prev.Next = nil
		} else {
			prev.Next = prev.Next.Next
		}
	}
}
