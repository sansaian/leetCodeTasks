package main

import (
	"fmt"

	"github.com/leetCodeTasks/datastruct/list"
)

func main() {
	//myList := list.New([]int{1, 2, 2, 1})
	//fmt.Println(isPalindrome(myList))
	//
	//myList = list.New([]int{1, 2})
	//fmt.Println(isPalindrome(myList))
	//
	//myList = list.New([]int{1})
	//fmt.Println(isPalindrome(myList))

	myList1 := list.New([]int{1, 2, 2, 1})
	fmt.Println(isPalindrome1(myList1))
	//myList := reverseList(myList1)
	//myList.Print()
	//fmt.Println(compare(myList, myList))
	//fmt.Println(findMiddle(myList))
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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome1(head *list.ListNode) bool {
	sublist := findMiddle(head)
	sublist = reverseList(sublist)
	return compare(sublist, head)
}

func findMiddle(head *list.ListNode) *list.ListNode {

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *list.ListNode) *list.ListNode {
	var prev *list.ListNode
	curr := head
	for curr != nil {
		nextNode := curr.Next
		curr.Next, prev, curr = prev, curr, nextNode
	}
	return prev
}

func compare(a, b *list.ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return true
}
