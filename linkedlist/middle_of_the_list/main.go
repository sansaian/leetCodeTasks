package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var fast, slow, midle int
	result := head
	for head != nil {
		head = head.Next
		fast++
		midle = fast / 2
		for slow < midle {
			slow++
			result = result.Next
		}

	}
	return result
}
