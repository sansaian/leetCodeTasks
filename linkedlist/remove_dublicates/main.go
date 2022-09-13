package main

import "github.com/leetCodeTasks/datastruct/list"

func main() {
	//exampleList := list.New([]int{5, 4, 4, 4, 3, 2, 1, 1})
	exampleList := list.New([]int{1, 1, 2, 3, 3})
	exampleList.Print()
	removeElements(exampleList)
	exampleList.Print()
}

func removeElements(head *list.ListNode) *list.ListNode {
	remove(head)
	return head
}

func remove(head *list.ListNode) {
	switch {
	case head == nil || head.Next == nil:
		return
	case head.Val == head.Next.Val:
		head.Val = head.Next.Val
		head.Next = head.Next.Next
		remove(head)
		return
	}
	remove(head.Next)
}
