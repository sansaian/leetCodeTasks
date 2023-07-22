package main

import (
	"github.com/leetCodeTasks/datastruct/list"
)

func main() {
	myList1 := list.New([]int{1, 2, 3, 4, 5, 6, 7, 8})
	myList1.Print()
	myList1 = reverseBetween(myList1, 2, 5)
	myList1.Print()
	myList1 = list.New([]int{1, 2, 3, 4, 5, 6, 7, 8})
	myList1.Print()
	myList1 = reverseBetween(myList1, 1, 2)
	myList1.Print()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *list.ListNode, left int, right int) *list.ListNode {
	// Return nil if the list is empty
	if head == nil {
		return nil
	}

	// Initialize a dummy node that points to the head
	dummy := &list.ListNode{Next: head}
	prev := dummy

	// Move 'prev' to its place which is at position 'left - 1'
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// Reverse the sublist from position 'left' to 'right'
	prevRev, currRev := reverse(prev.Next, right-left+1)

	// Connect the 'prev.Next' with the reversed sublist and
	// also connect the end of reversed sublist with the remaining list 'currRev'
	prev.Next.Next = currRev
	prev.Next = prevRev

	// Return the next node of dummy node
	return dummy.Next
}

func reverse(head *list.ListNode, len int) (*list.ListNode, *list.ListNode) {
	var prev *list.ListNode
	// Reverse 'len' nodes of the list and keep track of the next node after the reversed sublist
	for i := 0; i < len; i++ {
		nextRev := head.Next
		head.Next = prev
		prev = head
		head = nextRev
	}
	// Return the head of reversed sublist and the next node after the reversed sublist
	return prev, head
}
