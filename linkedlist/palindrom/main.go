package main

import (
	"fmt"
	"github.com/leetCodeTasks/datastruct/list"
)

func main() {
	myList := list.New([]int{1, 2, 2, 1})
	fmt.Println(isPalindrome(myList))

	myList = list.New([]int{1, 2})
	fmt.Println(isPalindrome(myList))

	myList = list.New([]int{1})
	fmt.Println(isPalindrome(myList))
}

func isPalindrome(head *list.ListNode) bool {
	var directly, revert []int
	if head == nil {
		return false
	}
	for head.Next != nil {
		directly = append(directly, head.Val)
		revert = append(revert, head.Val)
		head = head.Next
	}
	directly = append(directly, head.Val)
	revert = append(revert, head.Val)
	if len(directly) != len(revert) {
		return false
	}
	for i, v := range directly {
		if v != revert[len(revert)-1-i] {
			return false
		}
	}
	return true
}
