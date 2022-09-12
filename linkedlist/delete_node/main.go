package main

import "github.com/leetCodeTasks/datastruct/list"

func main() {
	//exampleList := list.New([]int{1, 2, 6, 3, 4, 6, 6})
	//exampleList := list.New([]int{7, 7, 7, 7})
	exampleList := list.New([]int{5, 4, 3, 2, 1, 1})

	exampleList.Print()
	removeElements(exampleList, 1)
	exampleList.Print()
}

func removeElements(head *list.ListNode, val int) *list.ListNode {
	remove(head, val)
	if head != nil && head.Val == val {
		head = nil
	}
	return head
}

func remove(head *list.ListNode, val int) {
	switch {
	case head == nil || head.Next == nil:
		return
	case head.Val == val:
		head.Val = head.Next.Val
		head.Next = head.Next.Next
		remove(head, val)
		return
	case head.Next.Val == val:
		head.Next = head.Next.Next
		remove(head, val)
	}
	remove(head.Next, val)
}
