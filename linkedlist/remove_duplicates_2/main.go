package main

import "github.com/leetCodeTasks/datastruct/list"

func main() {
	//myList := list.New([]int{1, 2, 3, 3, 4, 4, 5})
	//myList.Print()
	//deleteDuplicates(myList)
	//myList.Print()
	//
	//myList = list.New([]int{1, 2, 3, 3, 3, 4, 4, 5})
	//myList.Print()
	//deleteDuplicates(myList)
	//myList.Print()
	//
	//myList = list.New([]int{1, 1, 1, 2, 3}) //2,3
	//myList.Print()
	//myList = deleteDuplicates(myList)
	//myList.Print()
	//
	//myList = list.New([]int{1, 1}) //[]
	//myList.Print()
	//myList = deleteDuplicates(myList)
	//myList.Print()
	//
	//myList = list.New([]int{1, 1, 1}) //[]
	//myList.Print()
	//myList = deleteDuplicates(myList)
	//myList.Print()

	myList := list.New([]int{1, 2, 2, 2}) // [1,2,2,2] [1]
	myList.Print()
	myList = deleteDuplicates(myList)
	myList.Print()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *list.ListNode) *list.ListNode {
	dummy := &list.ListNode{Next: head}
	prev := dummy

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			prev.Next = head.Next
		} else {
			prev = prev.Next
		}
		head = head.Next
	}

	return dummy.Next
}
