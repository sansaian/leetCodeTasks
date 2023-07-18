package main

import "github.com/leetCodeTasks/datastruct/list"

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *list.ListNode) *list.ListNode {
	// Check edge case: linked list has 0 or 1 nodes, just return
	if head == nil || head.Next == nil {
		return head
	}
	dummy := head.Next      // Step 5
	var prev *list.ListNode // Initialize for step 3
	for head != nil && head.Next != nil {
		if prev != nil {
			prev.Next = head.Next // Step 4
		}
		prev = head // Step 3

		// Step 2
		nextNode := head.Next.Next
		head.Next.Next = head // Step 1

		head.Next = nextNode // Step 6
		head = nextNode      // Move to next pair (Step 3)
	}

	return dummy
}
