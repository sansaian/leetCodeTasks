package main

import (
	"github.com/leetCodeTasks/datastruct/list"
)

func mergeTwoLists(list1 *list.ListNode, list2 *list.ListNode) *list.ListNode {
	result := list1

	if list1 == nil {
		return list2
	}

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			tmp := &list.ListNode{
				Val:  list1.Val,
				Next: list1.Next,
			}
			list1.Val, list1.Next = list2.Val, tmp
			list2 = list2.Next
		}
		if list1.Next == nil {
			break
		}
		list1 = list1.Next
	}

	if list2 != nil {
		list1.Next = list2
	}
	return result
}

func main() {
	list1 := list.New([]int{1, 2, 4})
	list2 := list.New([]int{1, 3, 4})
	//exampleList.Print()
	result := mergeTwoLists(list1, list2)
	result.Print()
}
